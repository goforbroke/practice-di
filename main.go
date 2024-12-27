package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goforbroke/practice-di/handler"
	di "github.com/goforbroke/practice-di/internal"
)

func main() {
	c := di.NewContainer()
	weatherHandler := handler.NewWeatherHandler(c.NewUsecaseQuery())

	r := gin.Default()

	r.GET("/", weatherHandler.Weather)
	r.Run(":3000")
}
