package model

import (
	"encoding/base64"
	"math/big"
)

type Round struct {
	Address      string `json:"address"`
	Deposit      string `json:"deposit"`
	BalancesSnap string `json:"bsnap"`
	ParamsSnap   string `json:"psnap"`
	Spos         string `json:"Spos"`
	Sneg         string `json:"Sneg"`
	Reserve      string `json:"reserve"`
}

type PendingPlayer struct {
	Sender       string `json:"sender"`
	RoundAddress string `json:"roundAddress"`
}

type Player struct {
	Address      string `json:"address"`
	RoundAddress string `json:"roundAddress"`
	Balance      string `json:"balance"`
	Nwin         int64  `json:"nwin"`
	N            int64  `json:"n"`
	Spos         string `json:"spos"`
	Sneg         string `json:"sneg"`
}

type Lot struct {
	Address       string `json:"address"`
	RoundAddress  string `json:"roundAddress"`
	TimeFirst     string `json:"timeFirst"`
	TimeSecond    string `json:"timeSecond"`
	Value         string `json:"value"`
	Price         int64  `json:"price"`
	Owner         string `json:"owner"`
	ReceiveTokens string `json:"receiveTokens"`
	Snapshot      string `json:"snapshot"`
	PrevSnapshot  string `json:"prevSnapshot"`
}

func Base64ToInt(s string) (*big.Int, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	i := new(big.Int)
	i.SetBytes(data)
	return i, nil
}
