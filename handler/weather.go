package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goforbroke/practice-di/usecase"
)

type WeatherHandler interface {
	Weather(c *gin.Context)
}

type weatherHandler struct {
	usecaseQuery usecase.UsecaseQuery
}

func NewWeatherHandler(uq usecase.UsecaseQuery) WeatherHandler {
	return &weatherHandler{usecaseQuery: uq}
}

func (wh weatherHandler) Weather(c *gin.Context) {
	place := c.Query("place")
	w := wh.usecaseQuery.QueryToWeatherForecast(place)
	c.JSON(http.StatusOK, w)
}
