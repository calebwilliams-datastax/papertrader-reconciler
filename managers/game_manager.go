package managers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	m "github.com/calebwilliams-datastax/papertrader-api/models"
	util "github.com/calebwilliams-datastax/papertrader-api/util"
)

type GameManager struct {
	URL string
}

func NewGameManager(params map[string]string) GameManager {
	return GameManager{
		URL: fmt.Sprintf("%s/game", params["API_URL"]),
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

func (gm *GameManager) ReconcileGame(game m.Game) error {
	if time.Now().Unix() > game.End.Unix() {
		game.Finalized = true
		data, err := json.Marshal(game)
		if err != nil {
			return err
		}
		res, err := http.Post(fmt.Sprintf("%s/%s", gm.URL, game.ID), "application/json", strings.NewReader(string(data)))
		if err != nil {
			return err
		}
		if res.StatusCode != http.StatusOK {
			return fmt.Errorf("status code: %v returned from %s/%s", res.StatusCode, gm.URL, game.ID)
		}
	}
	return nil
}
