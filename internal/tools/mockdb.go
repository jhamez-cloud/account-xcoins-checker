package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex":{
		AuthToken: "123ABC",
		Username: "alex",
	},
	"bob":{
		AuthToken: "456DEF",
		Username: "bob",
	},
	"marie":{
		AuthToken: "789GHI",
		Username: "marie",
	},
}

var mockCoinBalance = map[string]CoinBalance{
	"alex":{
		Coins: 100,
		Username: "alex",
	},
	"bob":{
		Coins: 200,
		Username: "bob",
	},
	"marie":{
		Coins: 300,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails{
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok{
		return nil
	}

	return &clientData
}

func (d *mockDB) GetxCoinBalance(username string) *CoinBalance{
	time.Sleep(time.Second * 1)

	var clientData = CoinBalance{}
	clientData, ok := mockCoinBalance[username]
	if !ok{
		return nil
	}

	return &clientData
}

func(d *mockDB) SetupDatabase() error{
	return nil
}