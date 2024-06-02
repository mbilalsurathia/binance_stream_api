package service

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"marketdataservice/model"
	impl "marketdataservice/pkg/stream"
	"marketdataservice/serializer"
	"net/http"
	"strconv"
	"time"
)

// Binance API endpoints
const (
	apiBaseURL          = "https://api.binance.com"
	apiKlinesEndpoint   = "/api/v3/klines"
	binanceWebSocketURL = "wss://stream.binance.com:9443/ws/%s@kline_%s"
)

var (
	upgrade = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Binance struct {
}

func (b *Binance) Stream(c *gin.Context) serializer.Response {

	var (
		ReadingConn *websocket.Conn
		err         error
		conn1       *impl.Connection
	)
	queryParams := c.Request.URL.Query()

	var symbol string   //"btcusdt" //this can fetch from request
	var interval string //"1s"    //this can fetch from request
	_, ok := queryParams["symbol"]
	if ok {
		symbol = queryParams["symbol"][0]
	} else {
		return serializer.Response{
			Code:    400,
			Data:    true,
			Error:   "please provide the symbol in query params",
			Message: "please provide the symbol in query params",
		}
	}
	_, ok = queryParams["interval"]
	if ok {
		interval = queryParams["interval"][0]
	} else {
		return serializer.Response{
			Code:    400,
			Data:    true,
			Error:   "please provide the interval in query params",
			Message: "please provide the interval in query params",
		}
	}

	if ReadingConn, err = upgrade.Upgrade(c.Writer, c.Request, nil); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"1error": err.Error()})
	}

	if conn1, err = impl.InitConnection(ReadingConn); err != nil {
		ReadingConn.Close()
		c.JSON(http.StatusInternalServerError, gin.H{"1error": err.Error()})
	}
	// Open WebSocket connection for Binance stream
	url := fmt.Sprintf(binanceWebSocketURL, symbol, interval)
	fmt.Printf("url value %v\n", url)
	wsConn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		ReadingConn.Close() // Close the WebSocket connection
		conn1.Close()
		return serializer.Response{
			Code:    500,
			Data:    true,
			Error:   "please provide the interval in query params",
			Message: "please provide the interval in query params",
		}
	}
	defer wsConn.Close()
	defer ReadingConn.Close()
	defer conn1.Close()
	for {
		select {
		default:
			// Read message from Binance stream
			_, message, err := wsConn.ReadMessage()
			if err != nil {
				fmt.Printf("error reading message from binance err %v\n", err)
			}
			if err = conn1.WriteMessage(message); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"1error": err.Error()})
			}
		}
	}

}

// fetchHistorical fetches all historical trading data from Binance API
func (b *Binance) FetchHistorical(symbol string, interval string, startTimeQuery string, endTime time.Time, limit int) ([]model.Candlestick, error) {
	var allTrades []model.Candlestick
	var lastEndTime time.Time
	var startTime time.Time
	var err error
	if startTimeQuery != "" {
		layout := "2006-01-02"
		// Parse the user input into a time.Time object
		startTime, err = time.Parse(layout, startTimeQuery)
		if err != nil {
			fmt.Println("Error parsing input date:", err)
			return nil, err
		}
	} else {
		startTime = time.Now().AddDate(0, 0, -7) // 7 days ago
	}
	for {
		// Calculate the end time for this batch
		if lastEndTime.IsZero() {
			lastEndTime = endTime
		} else {
			lastEndTime = lastEndTime.Add(-time.Second * time.Duration(1)) // Adjust by 1 second to prevent duplicate records
		}

		// Fetch data for this batch
		trades, err := FetchHistoricalDataBatch(symbol, interval, startTime, lastEndTime, limit)
		if err != nil {
			return nil, err
		}
		allTrades = append(allTrades, trades...)

		// Check if we fetched all available data
		if len(trades) < 1000 {
			break
		}
		startTime = trades[len(trades)-1].Timestamp.Add(time.Second * time.Duration(1)) // Increment by 1 second to prevent overlapping records
	}

	return allTrades, nil
}

// fetchHistoricalDataBatch fetches a batch of historical trading data from Binance API
func FetchHistoricalDataBatch(symbol string, interval string, startTime time.Time, endTime time.Time, limit int) ([]model.Candlestick, error) {
	endpoint := fmt.Sprintf("%s%s?symbol=%s&interval=%s&startTime=%d&endTime=%d", apiBaseURL, apiKlinesEndpoint, symbol, interval, startTime.Unix()*1000, endTime.Unix()*1000)
	if limit > 0 {
		endpoint = endpoint + fmt.Sprintf("&limit=%d", limit)
	}
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data [][]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var trades []model.Candlestick
	for _, item := range data {
		open := item[1].(string)
		high := item[2].(string)
		low := item[3].(string)
		price := item[4].(string)
		quantity := item[5].(string)
		closeTime := int64(item[6].(float64))

		timestamp := time.Unix(int64(item[0].(float64))/1000, 0)
		openFloat, _ := strconv.ParseFloat(open, 64)
		highFloat, _ := strconv.ParseFloat(high, 64)
		lowFloat, _ := strconv.ParseFloat(low, 64)
		priceFloat, _ := strconv.ParseFloat(price, 64)
		quantityFloat, _ := strconv.ParseFloat(quantity, 64)

		trades = append(trades, model.Candlestick{
			Timestamp: timestamp,
			Open:      openFloat,
			High:      highFloat,
			Low:       lowFloat,
			Close:     priceFloat,
			Volume:    quantityFloat,
			CloseTime: closeTime,
		})
	}
	return trades, nil
}
