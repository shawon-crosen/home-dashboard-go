package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shawon-crosen/dashboard-go/weather"
)

func setRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	// Create API route group
	api := router.Group("/api")
	{
		// Add /hello GET route to router and define route handler function
		api.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"msg": "world"})
		})

		api.GET("/weather", func(ctx *gin.Context) {
			w := weather.Weather{Client: http.Client{}, Params: weather.NewForecastParams()}
			forecast := w.GetData()
			ctx.JSON(200, forecast)
		})
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
