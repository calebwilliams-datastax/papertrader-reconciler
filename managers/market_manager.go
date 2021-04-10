package managers

import (
	"fmt"
	// m "github.com/calebwilliams-datastax/papertrader-api/models"
	// util "github.com/calebwilliams-datastax/papertrader-api/util"
)

type MarketManager struct {
	URL string
}

func NewMarketManager(params map[string]string) OrderManager {
	return OrderManager{
		URL: fmt.Sprintf("%s/orders", params["API_URL"]),
	}

}

func (om *MarketManager) FetchPrice() (float64, error) {

	return 0, nil
}
