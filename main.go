package main

import (
	"flag"
	"fmt"

	"github.com/calebwilliams-datastax/papertrader-reconciler/managers"
	"github.com/calebwilliams-datastax/papertrader-reconciler/reconciler"
)

func main() {
	params := processFlags()
	fmt.Printf("-=papertrader-reconciler\nparams:%v\n", len(params))
	gm := managers.NewGameManager(params)
	pm := managers.NewPortfolioManager(params)
	om := managers.NewOrderManager(params)
	reconciler.Reconcile(gm, pm, om)

	fmt.Printf("%s, %s, %s\n", gm.URL, pm.URL, om.URL)
}

func processFlags() map[string]string {
	flags := map[string]string{}
	api_url := flag.String("API_URL", "http://localhost:8084", "-API_URL=http://localhost:8084")
	db_url := flag.String("DB_URL", "http://localhost:8082", "-DB_URL=http://localhost:8082")
	auth_url := flag.String("AUTH_URL", "http://localhost:8081", "-AUTH_URL=http://localhost:8081")
	db_user := flag.String("DB_USER", "cassandra", "-DB_USER=cassandra")
	db_pass := flag.String("DB_PASS", "cassandra", "-DB_PASS=cassandra")
	av_url := flag.String("AV_URL", "https://www.alphavantage.co/", "AV_URL=https://www.alphavantage.co/")
	av_token := flag.String("AV_TOKEN", "HR9QB2RM5GWOO0IX", "-AV_TOKEN=foo")
	cmdline := flag.String("CMDLINE", "false", "-CMDLINE=false")

	flag.Parse()
	flags["API_URL"] = *api_url
	flags["DB_URL"] = *db_url
	flags["DB_USER"] = *db_user
	flags["DB_PASS"] = *db_pass
	flags["AUTH_URL"] = *auth_url
	flags["AV_URL"] = *av_url
	flags["AV_TOKEN"] = *av_token
	flags["CMDLINE"] = *cmdline
	return flags
}
