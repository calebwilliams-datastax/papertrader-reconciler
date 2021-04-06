package managers

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/calebwilliams-datastax/papertrader-api/models"
	util "github.com/calebwilliams-datastax/papertrader-api/util"
)

type GameManager struct {
	URL string
}

func NewGameManager(params map[string]string) GameManager {
	return GameManager{
		URL: fmt.Sprintf("%s/games", params["API_URL"]),
	}

}

func (gm *GameManager) FetchOpenGames() (m.APIGameResponse, error) {
	games := m.APIGameResponse{}
	res, err := http.Get(fmt.Sprintf("%s/list/open", gm.URL))
	if err != nil {
		return games, err
	}
	data, err := util.ReadResponse(res)
	if err != nil {
		return games, err
	}
	json.Unmarshal([]byte(data), &games)
	return games, nil
}
