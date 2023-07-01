package polygon_io_demo

import (
	"context"
	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
	"log"
	"strconv"
	"testing"
	"time"
)

func Test_StockDaily(t *testing.T) {

	cli := polygon.New("LtsCZadfcYZn3qtiOQ8_3oDpnL7NJ3rQ")

	params := &models.GetGroupedDailyAggsParams{
		Locale:     models.US,
		MarketType: models.Stocks,
		Date:       models.Date(time.Date(2023, 6, 29, 0, 0, 0, 0, time.UTC)),
	}

	resp, err := cli.GetGroupedDailyAggs(context.Background(), params,
		models.APIKey("LtsCZadfcYZn3qtiOQ8_3oDpnL7NJ3rQ"),
		models.QueryParam("adjusted", strconv.FormatBool(true)))
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(resp)
}
