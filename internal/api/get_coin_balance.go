package api

import (
	"net/http"
	"encoding/json"
	"xs-coin-checker/internal/httpx"
	"xs-coin-checker/internal/tools"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetxCoinBalance(res http.ResponseWriter , req *http.Request){
	var decoder = schema.NewDecoder()
	var coinBalanceParams = httpx.CoinBalanceParams{}
	var err error

	err = decoder.Decode(&coinBalanceParams,req.URL.Query())
	if err != nil{
		log.Error(err)
		httpx.RequestErrorHandler(res,err)
		return
	}

	var database  *tools.DatabaseInterface
	database , err = tools.NewDatabase()
	if err != nil{
		httpx.InternalErrorHandler(res)
		return
	}

	var tokenDetails *tools.CoinBalance
	tokenDetails = (*database).GetxCoinBalance(coinBalanceParams.Username)

	if tokenDetails == nil{
		log.Error(err)
		httpx.InternalErrorHandler(res)
		return
	}

	var response = httpx.CoinBalanceParams{
		Username: coinBalanceParams.Username,
		Coins: tokenDetails.Coins,
	}

	res.Header().Set("Content-Type","Application/json")
	err = json.NewEncoder(res).Encode(response)
	if err != nil{
		log.Error(err)
		httpx.InternalErrorHandler(res)
		return
	}
}