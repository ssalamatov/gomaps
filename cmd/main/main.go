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
	router.POST("/city", server.CreateCityHandler)
	router.GET("/city/info", server.GetCityInfoHandler)
	router.GET("/city/:id", server.GetCountryHandler)
	router.DELETE("/city/:id", server.RemoveCityHandler)
	router.GET("/country/:id", server.GetCountryHandler)
	router.DELETE("/country/:id", server.RemoveCountryHandler)

	log.Fatal(http.ListenAndServe(config.GetAppAddr(), router))
}
