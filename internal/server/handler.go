package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ssalamatov/gomaps/internal/city"
	"github.com/ssalamatov/gomaps/pkg/client/postgresql"
)

func HandleErrorResponse(c *gin.Context, err error) {
	resp := make(map[string]string)
	if e := errors.Unwrap(err); e != nil {
		resp["error"] = e.Error()
	} else {
		resp["error"] = err.Error()
	}
	log.Println(err)
	c.JSON(http.StatusBadRequest, resp)
}

func GetValidateId(c *gin.Context) (int, error) {
	cid := c.Params.ByName("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		return 0, fmt.Errorf("id can be only integer. got %v: %w", cid, ErrValidation)
	}
	return id, nil
}

func GetValidateCity(c *gin.Context) (*city.CreateCityDTO, error) {
	var city city.CreateCityDTO

	if err := c.BindJSON(&city); err != nil {
		return nil, fmt.Errorf("invalid json body. got %v: %w", city, ErrDecodeBody)
	}
	return &city, nil
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

func (server *Server) CreateCityHandler(c *gin.Context) {
	city, err := GetValidateCity(c)
	if err != nil {
		HandleErrorResponse(c, err)
		return
	}

	err = postgresql.CreateCity(server.ctx, server.pool, city)
	if err != nil {
		HandleErrorResponse(c, err)
		return
	}
	c.Status(http.StatusNoContent)
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
