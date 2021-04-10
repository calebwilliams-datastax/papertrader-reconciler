package managers

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/calebwilliams-datastax/papertrader-api/models"
	util "github.com/calebwilliams-datastax/papertrader-api/util"
)

type PortfolioManager struct {
	URL string
}

func NewPortfolioManager(params map[string]string) PortfolioManager {
	return PortfolioManager{
		URL: fmt.Sprintf("%s/portfolio", params["API_URL"]),
	}

}

func (gm *PortfolioManager) FetchPortfoliosByGameID(gameID string) (m.APIPortfolioResponse, error) {
	portfolios := m.APIPortfolioResponse{}
	res, err := http.Get(fmt.Sprintf("%s/%s", gm.URL, gameID))
	if err != nil {
		return portfolios, err
	}
	data, err := util.ReadResponse(res)
	if err != nil {
		return portfolios, err
	}
	json.Unmarshal([]byte(data), &portfolios)
	return portfolios, nil
}
