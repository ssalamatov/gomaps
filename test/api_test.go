package api_test

import (
	"net/http"
	"strings"
	"testing"

	"github.com/ssalamatov/gomaps/test/client"
)

const URL = "http://127.0.0.1:8080"

func GetFullUrl(resourse string) string {
	return strings.Join([]string{URL, resourse}, "/")
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected code %d. Got %d\n", expected, actual)
	}
}

func CheckResponseContent(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected content \n%s. Got \n%s\n", expected, actual)
	}
}

func CheckResponse(t *testing.T, ExpCode int, ExpContent string, resp *client.Response) {
	CheckResponseCode(t, ExpCode, resp.Code)
	CheckResponseContent(t, ExpContent, string(resp.Content))
}

func TestAPIGetEmptyCities(t *testing.T) {
	// Test empty response when no cities exist
	client.PrepareDB(false)

	req, err := http.NewRequest("GET", GetFullUrl("cities"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusOK,
		`[]`,
		resp)
}

func TestAPIGetCities(t *testing.T) {
	client.PrepareDB(true)

	req, err := http.NewRequest("GET", GetFullUrl("cities"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusOK,
		`[{"id":1,"name":"Moscow","is_capital":true,"population":5,"found_at":"2022-03-19T23:36:13.183732Z"}]`,
		resp)
}

func TestAPIGetCityInfo(t *testing.T) {
	client.PrepareDB(true)

	req, err := http.NewRequest("GET", GetFullUrl("city/info?name=Moscow"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusOK,
		`[{"id":1,"name":"Moscow","is_capital":true,"population":5,"found_at":"2022-03-19T23:36:13.183732Z","country":"Russia"}]`,
		resp)
}

func TestAPIGetCityInfoNotFound(t *testing.T) {
	// Invalid city not found
	client.PrepareDB(true)

	req, err := http.NewRequest("GET", GetFullUrl("city/info?name=A"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusOK,
		`[]`,
		resp)
}

func TestAPIGetCountries(t *testing.T) {
	client.PrepareDB(true)

	req, err := http.NewRequest("GET", GetFullUrl("countries"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusOK,
		`[{"id":1,"name":"Russia"},{"id":2,"name":"USA"}]`,
		resp)
}

func TestAPIGetEmptyCountries(t *testing.T) {
	client.PrepareDB(false)

	req, err := http.NewRequest("GET", GetFullUrl("countries"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusOK,
		`[]`,
		resp)
}

func TestAPIGetCountryById(t *testing.T) {
	client.PrepareDB(true)

	req, err := http.NewRequest("GET", GetFullUrl("country/1"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusOK,
		`{"id":1,"name":"Russia"}`,
		resp)
}

func TestAPIGetCountryByIdNotFound(t *testing.T) {
	client.PrepareDB(true)

	req, err := http.NewRequest("GET", GetFullUrl("country/5"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusBadRequest,
		`{"error":"not found"}`,
		resp)
}

func TestAPIGetCountryByIdError(t *testing.T) {
	client.PrepareDB(true)

	req, err := http.NewRequest("GET", GetFullUrl("country/a"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusBadRequest,
		`{"error":"validation failed"}`,
		resp)
}

func TestAPICreateCity(t *testing.T) {
	client.PrepareDB(true)

	city := strings.NewReader(`{"name":"Ufa","is_capital":false,"population":1,"found_at":"2022-03-19T23:36:13.183732Z","country":"Russia"}`)
	req, err := http.NewRequest("POST", GetFullUrl("city"), city)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusNoContent,
		``,
		resp)
}

func TestAPICreateCityCountryNotFound(t *testing.T) {
	client.PrepareDB(false)

	city := strings.NewReader(`{"name":"Ufa","is_capital":false,"population":1,"found_at":"2022-03-19T23:36:13.183732Z","country":"Russia"}`)
	req, err := http.NewRequest("POST", GetFullUrl("city"), city)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusBadRequest,
		`{"error":"not found"}`,
		resp)
}

func TestAPICreateCityValidationFailed(t *testing.T) {
	client.PrepareDB(true)

	data := []string{
		`{"is_capital":false,"population":1,"found_at":"2022-03-19T23:36:13.183732Z","country":"Russia"}`,
		`{"name":"Ufa","is_capital":false,"found_at":"2022-03-19T23:36:13.183732Z","country":"Russia"}`,
		`{"name":"Ufa","is_capital":false,"population":1,"country":"Russia"}`,
		`{"name":"Ufa","is_capital":false,"population":1,"found_at":"2022-03-19T23:36:13.183732Z"}`}

	for _, buf := range data {
		city := strings.NewReader(buf)
		req, err := http.NewRequest("POST", GetFullUrl("city"), city)
		if err != nil {
			t.Errorf("Failed request %v", err)
		}

		resp := client.NewClient().Execute(req)
		CheckResponse(
			t,
			http.StatusBadRequest,
			`{"error":"body decode failed"}`,
			resp)
	}
}

func TestAPICreateCityOptionalFields(t *testing.T) {
	// is_capiatal is optioanl value. can be omitted.
	client.PrepareDB(true)

	data := []string{
		`{"name":"Ufa","population":1,"found_at":"2022-03-19T23:36:13.183732Z","country":"Russia"}`}

	for _, buf := range data {
		city := strings.NewReader(buf)
		req, err := http.NewRequest("POST", GetFullUrl("city"), city)
		if err != nil {
			t.Errorf("Failed request %v", err)
		}

		resp := client.NewClient().Execute(req)
		CheckResponse(
			t,
			http.StatusNoContent,
			``,
			resp)
	}
}

func TestAPIRemoveCity(t *testing.T) {
	client.PrepareDB(true)

	req, err := http.NewRequest("DELETE", GetFullUrl("city/1"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusNoContent,
		``,
		resp)
}

func TestAPIRemoveCityNotFound(t *testing.T) {
	client.PrepareDB(false)

	req, err := http.NewRequest("DELETE", GetFullUrl("city/1"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusBadRequest,
		`{"error":"not found"}`,
		resp)
}

func TestAPIRemoveCountryNotFound(t *testing.T) {
	client.PrepareDB(false)

	req, err := http.NewRequest("DELETE", GetFullUrl("country/1"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusBadRequest,
		`{"error":"not found"}`,
		resp)
}

func TestAPIRemoveCountry(t *testing.T) {
	client.PrepareDB(true)

	req, err := http.NewRequest("DELETE", GetFullUrl("country/2"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusNoContent,
		``,
		resp)
}

func TestAPIRemoveCountryError(t *testing.T) {
	// Countries with linked cities can not be removed
	client.PrepareDB(true)

	req, err := http.NewRequest("DELETE", GetFullUrl("country/1"), nil)
	if err != nil {
		t.Errorf("Failed request %v", err)
	}

	resp := client.NewClient().Execute(req)
	CheckResponse(
		t,
		http.StatusBadRequest,
		`{"error":"sql failed query"}`,
		resp)
}
