package tools

import (
	"time"
)

type mockDB struct {}

var mockLoginDetails = map[string]LoginDetails {
	"alex": {
		AuthToken: "123ABC",
		Username: "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username: "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username: "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails {
	"alex": {
        Coins: 1000,
        Username: "alex",
    },
    "jason": {
        Coins: 2000,
        Username: "jason",
    },
    "marie": {
        Coins: 3000,
        Username: "marie",
    },
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	//Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	//Simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
    // Mock setup, no actual database setup
    return nil
}