package model

import (
	"time"
)

type Candlestick struct {
	Timestamp                time.Time `json:"openTime"`
	Open                     float64   `json:"open"`
	High                     float64   `json:"high"`
	Low                      float64   `json:"low"`
	Close                    float64   `json:"close"`
	Volume                   float64   `json:"volume"`
	CloseTime                int64     `json:"closeTime"`
	QuoteAssetVolume         float64   `json:"quoteAssetVolume"`
	NumberOfTrades           int       `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume  float64   `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume float64   `json:"takerBuyQuoteAssetVolume"`
}

type HistoricalDataResponse struct {
	Data []Candlestick
}
