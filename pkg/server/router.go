package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shawon-crosen/dashboard-go/pkg/config"
	"github.com/shawon-crosen/dashboard-go/pkg/cta"
	"github.com/shawon-crosen/dashboard-go/pkg/weather"
)

func setRouter(conf []byte) *gin.Engine {
	// Creates default gin router with Logger and Recovery middleware already attached
	router := gin.Default()
	configData := config.GenerateConfig(conf)
	// Create API route group
	api := router.Group("/api")
	{
		api.GET("/weather", func(ctx *gin.Context) {
			w := weather.Weather{Client: http.Client{}, Params: weather.NewForecastParams(configData.WeatherConfig)}
			forecastResp := w.GetData()
			ctx.JSON(200, w.FormatData(forecastResp))
		})

		api.GET("/cta", func(ctx *gin.Context) {
			trains := cta.AllTrains{Client: http.Client{}, ApiKey: configData.CtaConfig.Api_key, Stations: configData.CtaConfig.Stations}
			tResp := trains.GetTrains()
			ctx.JSON(200, trains.FormatData(tResp))
		})
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
