package server

import (
	"marketdataservice/bootstraper"
	"marketdataservice/controller"
	"marketdataservice/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(mode string, b *bootstraper.Bootstrapper) *gin.Engine {
	gin.SetMode(mode)

	router := gin.Default()
	router.Use(middleware.Cors())
	router.Use(func(c *gin.Context) {
		c.Set("currency_pair", b.CurrencyPair) // Inject database connection into context
		c.Next()
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/stream", controller.Stream)
		v1.GET("/historical_data", controller.HistoricalData)
		v1.GET("/get_specific_historical_data", controller.GetSpecificHistoricalData)
		v1.GET("/get_currency_pairs_values", controller.GetCurrencyManagementValues)
	}

	return router
}
