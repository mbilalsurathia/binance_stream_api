package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"marketdataservice/model"
	"marketdataservice/serializer"
	"marketdataservice/service"
	"net/http"
	"time"
)

func Stream(c *gin.Context) {
	var binanceService service.Binance
	if err := c.Bind(&binanceService); err == nil {
		res := binanceService.Stream(c)
		if res.Code == 200 {
			c.JSON(http.StatusOK, res.Data)
		} else {
			c.JSON(res.Code, gin.H{"error": res.Message})
		}
	} else {
		c.JSON(http.StatusBadRequest, serializer.ParamError("Bad Input", err))
	}
}

func HistoricalData(c *gin.Context) {
	var binanceService service.Binance
	queryParams := c.Request.URL.Query()
	var symbol string         //"btcusdt" //this can fetch from request
	var interval string       //"1s"    //this can fetch from request
	var limit int             //"100"    //this can fetch from request
	var startTimeQuery string //"100"    //this can fetch from request
	_, ok := queryParams["symbol"]
	if ok {
		symbol = queryParams["symbol"][0]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please provide the symbol in query params"})
		return
	}
	_, ok = queryParams["interval"]
	if ok {
		interval = queryParams["interval"][0]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please provide the interval in query params"})
		return
	}
	_, ok = queryParams["limit"]
	if ok {
		limit = cast.ToInt(queryParams["limit"][0])
	}
	_, ok = queryParams["starttime"]
	if ok {
		startTimeQuery = cast.ToString(queryParams["starttime"][0])
	}
	endTime := time.Now()
	trades, err := binanceService.FetchHistorical(symbol, interval, startTimeQuery, endTime, limit)
	if err != nil {
		fmt.Println("Error fetching historical data:", err)
	} else {
		var res model.HistoricalDataResponse
		for _, trade := range trades {
			res.Data = append(res.Data, trade)
			fmt.Printf("%s - Price: %.2f, Quantity: %.2f\n", trade.Timestamp.Format(time.RFC3339), trade.Close, trade.Volume)
		}
		c.JSON(200, res)
	}
}

func GetSpecificHistoricalData(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	var symbol string         //"btcusdt" //this can fetch from request
	var interval string       //"1s"    //this can fetch from request
	var limit int             //"100"    //this can fetch from request
	var startTimeQuery string //"100"    //this can fetch from request
	_, ok := queryParams["symbol"]
	if ok {
		symbol = queryParams["symbol"][0]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please provide the symbol in query params"})
		return
	}
	_, ok = queryParams["interval"]
	if ok {
		interval = queryParams["interval"][0]
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please provide the interval in query params"})
		return
	}
	_, ok = queryParams["limit"]
	if ok {
		limit = cast.ToInt(queryParams["limit"][0])
	}
	_, ok = queryParams["starttime"]
	if ok {
		startTimeQuery = cast.ToString(queryParams["starttime"][0])
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "please provide the starttime in query params"})
		return
	}
	layout := "2006-01-02"
	// Parse the user input into a time.Time object
	startTime, err := time.Parse(layout, startTimeQuery)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "please provide the correct format for start time 2006-01-01"})
		return
	}
	//startTime := time.Now().AddDate(0, 0, -1) // 7 days ago
	endTime := time.Now()
	trades, err := service.FetchHistoricalDataBatch(symbol, interval, startTime, endTime, limit)
	if err != nil {
		fmt.Println("Error fetching historical data:", err)
	} else {
		fmt.Sprintf("H %v:", symbol)
		var res model.HistoricalDataResponse
		for _, trade := range trades {
			res.Data = append(res.Data, trade)
			fmt.Printf("%s - Price: %.2f, Quantity: %.2f\n", trade.Timestamp.Format(time.RFC3339), trade.Close, trade.Volume)
		}

		c.JSON(200, res)
	}
}

func GetCurrencyManagementValues(c *gin.Context) {
	db := c.MustGet("currency_pair").(model.CurrencyPair)
	data, err := db.GetAllCurrencyPairs()
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "some error occurred"})
		return
	}
	c.JSON(200, data)
}
