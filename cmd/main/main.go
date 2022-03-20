package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ssalamatov/gomaps/internal/config"
	"github.com/ssalamatov/gomaps/internal/server"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	server, err := server.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/cities", server.GetAllCitiesHandler)
	router.GET("/countries", server.GetAllCountriesHandler)
	router.GET("/countries/:id", server.GetCountryHandler)
	router.GET("/city/:id", server.GetCountryHandler)
	router.DELETE("/city/:id", server.RemoveCityHandler)
	router.GET("/city/info", server.GetCityInfoHandler)

	log.Fatal(http.ListenAndServe(config.GetAppAddr(), router))
}