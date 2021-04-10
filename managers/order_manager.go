package managers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/calebwilliams-datastax/papertrader-api/models"
	m "github.com/calebwilliams-datastax/papertrader-api/models"
	util "github.com/calebwilliams-datastax/papertrader-api/util"
)

type OrderManager struct {
	URL string
}

func NewOrderManager(params map[string]string) OrderManager {
	return OrderManager{
		URL: fmt.Sprintf("%s/order", params["API_URL"]),
	}

}

func (om *OrderManager) FetchOpenGames() (m.APIOrderResponse, error) {
	orders := m.APIOrderResponse{}
	res, err := http.Get(fmt.Sprintf("%s/list/open", om.URL))
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

func (om *OrderManager) FetchOrdersByPortfolioID(portfolioID string) (m.APIOrderResponse, error) {
	orders := m.APIOrderResponse{}
	res, err := http.Get(fmt.Sprintf("%s/portfolio/%s", om.URL, portfolioID))
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

func (om *OrderManager) ReconcileOrder(order m.Order, balance float64, gameEnd time.Time) (float64, error) {
	log.Printf("reconciling order: %s\n", order.ID)
	transactionCost := 0
	switch order.OrderAction {
	case models.Buy:
		return om.ReconcileBuyOrder(order, balance, gameEnd)
	case models.Sell:
		return om.ReconcileSellOrder(order, balance, gameEnd)
	}

	return float64(transactionCost), nil
}

func (om *OrderManager) ReconcileBuyOrder(order m.Order, balance float64, gameEnd time.Time) (float64, error) {
	return -9999, nil
}

func (om *OrderManager) ReconcileSellOrder(order m.Order, balance float64, gameEnd time.Time) (float64, error) {
	return 0, nil
}
