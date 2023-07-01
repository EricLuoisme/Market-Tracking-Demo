package polygon_io_demo

import (
	"context"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"log"
	"testing"
	"time"
)

var (
	apiKey   = ""
	adjusted = true
	cli      = polygon.New(apiKey)
)

// Test_StockDaily is for daily stock market place info
func Test_StockDaily(t *testing.T) {

	params := models.GetGroupedDailyAggsParams{
		Locale:     models.US,
		MarketType: models.Stocks,
		Date:       models.Date(time.Date(2023, 6, 29, 0, 0, 0, 0, time.UTC)),
	}.WithAdjusted(true)

	resp, err := cli.GetGroupedDailyAggs(context.Background(), params)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}

// Test_OpenCloseDaily is for daily open & close position
func Test_OpenCloseDaily(t *testing.T) {

	params := models.GetDailyOpenCloseAggParams{
		Ticker: "APPL",
		Date:   models.Date(time.Date(2020, 10, 14, 0, 0, 0, 0, time.Local)),
	}.WithAdjusted(true)

	agg, err := cli.GetDailyOpenCloseAgg(context.Background(), params)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(agg)
}
