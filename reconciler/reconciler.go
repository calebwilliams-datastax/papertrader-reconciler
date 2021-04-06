package reconciler

import (
	"time"

	models "github.com/calebwilliams-datastax/papertrader-api/models"
	"github.com/calebwilliams-datastax/papertrader-reconciler/managers"
)

type ReconcilerReport struct {
	ID          string    `json:"id"`
	Created     time.Time `json:"created"`
	GameID      string    `json:"game_id"`
	PortfolioID string    `json:"portfolio_id"`
	OrderID     string    `json:"order_id"`
	Note        string    `json:"note"`
	Error       error     `json:"error"`
}

func ReconcileGame(g models.Game, o managers.OrderManager, p managers.PortfolioManager) ([]ReconcilerReport, error) {
	report := []ReconcilerReport{}

	portfolioAPIRes, err := p.FetchPortfoliosByGameID(g.ID)
	if err != nil {
		return report, err
	}
	for _, portfolio := range portfolioAPIRes.Data {
		orderAPIRes, err := o.FetchOrdersByPortfolioID(portfolio.ID)
		if err != nil {
			return report, err
		}
		for _, order := range orderAPIRes.Data {
			r := ReconcilerReport{
				ID:          models.GenerateID(),
				Created:     time.Now(),
				GameID:      g.ID,
				PortfolioID: order.ID,
			}
			//o, err := ReconcileOrder(order)
			if err != nil {
				r.Error = err
				report = append(report, r)
				break
			}

			report = append(report, r)
		}
	}
	return report, nil
}
