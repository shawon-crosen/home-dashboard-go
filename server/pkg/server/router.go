package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shawon-crosen/dashboard-go/server/pkg/config"
	"github.com/shawon-crosen/dashboard-go/server/pkg/quotes"
	"github.com/shawon-crosen/dashboard-go/server/pkg/weather"
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
				hourly.GET("", func(ctx *gin.Context) {
					w := weather.Weather{Client: http.Client{}, Params: weather.NewForecastParams(configData.WeatherConfig)}
					forecastResp := w.GetData("hourly")
					if forecastResp != nil {
						ctx.JSON(200, w.FormatHourlyData(*forecastResp, "hourly"))
					} else {
						ctx.JSON(500, "A server error has occured")
					}

				})
			}

			daily := forecast.Group("/daily")
			{
				daily.GET("", func(ctx *gin.Context) {
					w := weather.Weather{Client: http.Client{}, Params: weather.NewForecastParams(configData.WeatherConfig)}
					forecastResp := w.GetData("daily")
					if forecastResp != nil {
						ctx.JSON(200, w.FormatDailyData(*forecastResp, "daily"))
					} else {
						ctx.JSON(500, "A server error has occured")
					}

				})
			}

			current := forecast.Group("/current")
			{
				current.GET("", func(ctx *gin.Context) {
					w := weather.Weather{Client: http.Client{}, Params: weather.NewForecastParams(configData.WeatherConfig)}
					forecastResp := w.GetData("current")
					if forecastResp != nil {
						ctx.JSON(200, w.FormatCurrentData(*forecastResp, "current"))
					} else {
						ctx.JSON(500, "A server error has occured")
					}

				})
			}
		}
		quote := api.Group("/quotes")
		{
			randQuote := quote.Group("random")
			{
				randQuote.GET("", func(ctx *gin.Context) {
					q := quotes.Quotes{Client: http.Client{}}
					qResp := q.QueryForQuote()
					if qResp != nil {
						ctx.JSON(200, qResp)
					} else {
						ctx.JSON(500, "A server error has occured")
					}
				})
			}
		}
	}

	router.NoRoute(func(ctx *gin.Context) { ctx.JSON(http.StatusNotFound, gin.H{}) })

	return router
}
