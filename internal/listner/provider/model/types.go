package model

import (
	"math/big"
)

type Round struct {
	Address      string   `json:"address"`
	Deposit      int      `json:"deposit"`
	BalancesSnap *big.Int `json:"bsnap"`
	ParamsSnap   *big.Int `json:"psnap"`
	Spos         *big.Int `json:"Spos"`
	Sneg         *big.Int `json:"Sneg"`
	Reserve      *big.Int `json:"reserve"`
}

type PendingPlayer struct {
	Sender       string `json:"sender"`
	RoundAddress string `json:"roundAddress"`
}

type Player struct {
	Address      string `json:"address"`
	RoundAddress string `json:"roundAddress"`
	Balance      int    `json:"balance"`
	Nwin         int    `json:"nwin"`
	N            int    `json:"n"`
	Spos         int    `json:"spos"`
	Sneg         int    `json:"sneg"`
}

type Lot struct {
	Address       string `json:"address"`
	RoundAddress  string `json:"roundAddress"`
	TimeFirst     int    `json:"timeF"`
	TimeSecond    int    `json:"timeS"`
	Value         int    `json:"value"`
	Price         int    `json:"price"`
	Owner         string `json:"owner"`
	ReceiveTokens int    `json:"receiveTokens"`
	Snapshot      int    `json:"snapshot"`
}
