package managers

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/calebwilliams-datastax/papertrader-api/models"
	util "github.com/calebwilliams-datastax/papertrader-api/util"
)

type OrderManager struct {
	URL string
}

func NewOrderManager(params map[string]string) OrderManager {
	return OrderManager{
		URL: fmt.Sprintf("%s/orders", params["API_URL"]),
	}

}

func (gm *OrderManager) FetchOpenGames() (m.APIOrderResponse, error) {
	orders := m.APIOrderResponse{}
	res, err := http.Get(fmt.Sprintf("%s/list/open", gm.URL))
	if err != nil {
		return orders, err
	}
	data, err := util.ReadResponse(res)
	if err != nil {
		return orders, err
	}
	json.Unmarshal([]byte(data), &orders)
	return orders, nil
}
