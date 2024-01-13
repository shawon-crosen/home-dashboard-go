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
		forecast := api.Group("/weather")
		{
			hourly := forecast.Group("/hourly")
			{
				hourly.GET("/", func(ctx *gin.Context) {
					w := weather.Weather{Client: http.Client{}, Params: weather.NewForecastParams(configData.WeatherConfig)}
					forecastResp := w.GetData("hourly")
					if forecastResp != nil {
						ctx.JSON(200, w.FormatData(*forecastResp, "hourly"))
					} else {
						ctx.JSON(500, "A server error has occured")
					}

				})
			}

			daily := forecast.Group("/daily")
			{
				daily.GET("/", func(ctx *gin.Context) {
					w := weather.Weather{Client: http.Client{}, Params: weather.NewForecastParams(configData.WeatherConfig)}
					forecastResp := w.GetData("daily")
					if forecastResp != nil {
						ctx.JSON(200, w.FormatData(*forecastResp, "daily"))
					} else {
						ctx.JSON(500, "A server error has occured")
					}

				})
			}

			current := forecast.Group("/current")
			{
				current.GET("/", func(ctx *gin.Context) {
					w := weather.Weather{Client: http.Client{}, Params: weather.NewForecastParams(configData.WeatherConfig)}
					forecastResp := w.GetData("current")
					if forecastResp != nil {
						ctx.JSON(200, w.FormatData(*forecastResp, "current"))
					} else {
						ctx.JSON(500, "A server error has occured")
					}

				})
			}
		}

		api.GET("/cta", func(ctx *gin.Context) {
			trains := cta.AllTrains{Client: http.Client{}, ApiKey: configData.CtaConfig.Api_key, Stations: configData.CtaConfig.Stations}
			tResp := trains.GetTrains()
			if tResp != nil {
				ctx.JSON(200, trains.FormatData(*tResp))
			} else {
				ctx.JSON(500, "A server error has occured")
			}
		})
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
