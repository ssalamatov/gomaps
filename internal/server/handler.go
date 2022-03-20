package server

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ssalamatov/gomaps/pkg/client/postgresql"
)

func HandleErrorResponse(c *gin.Context, err error) {
	response := make(map[string]string)
	response["error"] = err.Error()
	c.JSON(http.StatusBadRequest, response)
}

func GetValidateId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		return 0, NewErrValidation(err)
	}
	return id, nil
}

func (server *Server) GetAllCountriesHandler(c *gin.Context) {
	countries, err := postgresql.GetCountries(server.ctx, server.pool)
	if err != nil {
		HandleErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, countries)
}

func (server *Server) GetCountryHandler(c *gin.Context) {
	id, err := GetValidateId(c)
	if err != nil {
		HandleErrorResponse(c, err)
		return
	}
	country, err := postgresql.GetCountryById(server.ctx, server.pool, id)
	if err != nil {
		HandleErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, country)
}

func (server *Server) GetAllCitiesHandler(c *gin.Context) {
	cities, err := postgresql.GetCities(server.ctx, server.pool)
	if err != nil {
		HandleErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, cities)
}

func (server *Server) GetCityInfoHandler(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		c.String(http.StatusBadRequest, "Missing query parameter")
		return
	}

	city, err := postgresql.GetCityInfo(server.ctx, server.pool, name)
	if err != nil {
		HandleErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, city)
}

func (server *Server) RemoveCityHandler(c *gin.Context) {
	id, err := GetValidateId(c)
	if err != nil {
		HandleErrorResponse(c, err)
		return
	}

	err = postgresql.RemoveCity(server.ctx, server.pool, id)
	if err != nil {
		HandleErrorResponse(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func (server *Server) RemoveCountryHandler(c *gin.Context) {
	id, err := GetValidateId(c)
	if err != nil {
		HandleErrorResponse(c, err)
		return
	}

	err = postgresql.RemoveCountry(server.ctx, server.pool, id)
	if err != nil {
		HandleErrorResponse(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}
