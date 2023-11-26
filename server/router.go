package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shawon-crosen/dashboard-go/cta"
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
			forecastResp := w.GetData()
			ctx.JSON(200, w.FormatData(forecastResp))
		})

		api.GET("/cta", func(ctx *gin.Context) {
			t := cta.Station{Client: http.Client{}, Id: 41340}
			tResp := t.GetTrains()
			ctx.JSON(200, tResp)
		})
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
