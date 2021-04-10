package reconciler

import (
	"fmt"
	"log"
	"strconv"
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

func NewReport() ReconcilerReport {
	return ReconcilerReport{
		ID:      models.GenerateID(),
		Created: time.Now(),
	}
}

func (r *ReconcilerReport) WithGameID(gameID string) *ReconcilerReport {
	r.GameID = gameID
	return r
}

func (r *ReconcilerReport) WithOrderID(orderID string) *ReconcilerReport {
	r.OrderID = orderID
	return r
}

func (r *ReconcilerReport) WithPortfolioID(portfolioID string) *ReconcilerReport {
	r.PortfolioID = portfolioID
	return r
}

func (r *ReconcilerReport) WithError(err error) *ReconcilerReport {
	r.Error = err
	return r
}

func (r *ReconcilerReport) WithNote(note string) *ReconcilerReport {
	r.Note = note
	return r
}

func Reconcile(gm managers.GameManager, pm managers.PortfolioManager, om managers.OrderManager) ([]ReconcilerReport, error) {
	report := []ReconcilerReport{}

	games, err := gm.FetchOpenGames()
	if err != nil {
		log.Fatalf("could not fetch open games")
		return report, err
	}
	for _, game := range games.Data {
		r := NewReport()
		r.WithGameID(game.ID)
		fmt.Printf("reconciling game: %s\n", game.ID)
		if err := gm.ReconcileGame(game); err != nil {
			r.WithError(err).WithNote("failed to finalized game")
			report = append(report, r)
			break
		}
		portfolios, err := pm.FetchPortfoliosByGameID(game.ID)
		if err != nil {
			r.WithError(err).WithNote("failed fetching portfolios")
			report = append(report, r)
			break
		}
		//portfolios
		for _, portfolio := range portfolios.Data {
			r.WithPortfolioID(portfolio.ID)
			balance, err := strconv.ParseFloat(portfolio.Cash, 64)
			if err != nil {
				r.WithError(err).WithNote("could not parse portfolio cash balance")
				report = append(report, r)
				break
			}
			orders, err := om.FetchOrdersByPortfolioID(portfolio.ID)
			if err != nil {
				r.WithError(err).WithNote("could not fetch orders")
				report = append(report, r)
				break
			}
			//orders
			for _, order := range orders.Data {
				r.WithOrderID(order.ID)
				cost, err := om.ReconcileOrder(order, balance, game.End)
				if err != nil {
					r.WithError(err).WithNote("could not reconcile order")
					report = append(report, r)
					break
				}
				balance += cost
				r.WithNote(fmt.Sprintf("new cash balance: %v", balance))
				report = append(report, r)
				//alphavantage price here?
			}
		}
	}
	return report, nil
}

func ReconcileOrder(g models.Game, o models.Order, om managers.OrderManager, p managers.PortfolioManager) (ReconcilerReport, error) {
	r := ReconcilerReport{
		ID:          models.GenerateID(),
		Created:     time.Now(),
		GameID:      g.ID,
		PortfolioID: o.ID,
	}
	switch o.OrderAction {
	case models.Buy:
	}
	return r, nil
}
