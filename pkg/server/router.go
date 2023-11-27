package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shawon-crosen/dashboard-go/pkg/cta"
	"github.com/shawon-crosen/dashboard-go/pkg/weather"
)

func setRouter() *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()

	// Create API route group
	api := router.Group("/api")
	{
		api.GET("/weather", func(ctx *gin.Context) {
			w := weather.Weather{Client: http.Client{}, Params: weather.NewForecastParams()}
			forecastResp := w.GetData()
			ctx.JSON(200, w.FormatData(forecastResp))
		})

		api.GET("/cta", func(ctx *gin.Context) {
			t := cta.Station{Client: http.Client{}, Id: 41400}
			tResp := t.GetTrains()
			ctx.JSON(200, t.FormatData(tResp))
		})
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
