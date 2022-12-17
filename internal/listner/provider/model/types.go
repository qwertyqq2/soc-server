package model

import (
	"encoding/base64"
	"encoding/json"
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

type RoundWrapper struct {
	Name string `json:"name"`
	Data *Round `json:"data"`
}

func NewRoundWrapper(round *Round) *RoundWrapper {
	return &RoundWrapper{
		Name: "round",
		Data: round,
	}
}

type LotWrapper struct {
	Name string `json:"name"`
	Data *Lot   `json:"data"`
}

func NewLotWrapper(lot *Lot) *LotWrapper {
	return &LotWrapper{
		Name: "lot",
		Data: lot,
	}
}

type PlayerWrapper struct {
	Name string  `json:"name"`
	Data *Player `json:"data"`
}

func NewPlayerWrapper(player *Player) *PlayerWrapper {
	return &PlayerWrapper{
		Name: "player",
		Data: player,
	}
}

type Resp struct {
	resp []interface{}
}

func NewResp(data ...interface{}) *Resp {
	return &Resp{
		resp: data,
	}
}

func (r *Resp) GetResp() []interface{} {
	return r.resp
}

func (resp *Resp) Marshal() ([]byte, error) {
	data := resp.GetResp()
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
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
