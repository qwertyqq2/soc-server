package store

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/qwertyqq2/soc-server/internal/listner/provider/model"
)

type LotResp struct {
	Name string     `json:"name"`
	Data *model.Lot `json:"data"`
}

type PlayerResp struct {
	Name string        `json:"name"`
	Data *model.Player `json:"data"`
}

type Resp struct {
	resp []interface{}
}

func NewLotTest() *model.Lot {
	addr := "0xfB3Ee894f27d8d6798f0034d7763b7ad46C31B92"
	roundAddr := "0x5f63C25ee7f258706803E97d6a866cDe363201c2"
	owner := "0xFFFbEE9C752A1ca301A3821308F630609a54369A"
	timeF := "123345534234234"
	timeS := "133345534234234"
	val := "20"
	price := 60
	snap := "15468303776996484186685289508981412665698293446761833792330076793367195714725"
	prevSnap := "57003951752417535200923101159811040833146496380398724999470920633816778075019"
	return &model.Lot{
		Address:      addr,
		RoundAddress: roundAddr,
		Owner:        owner,
		TimeFirst:    timeF,
		TimeSecond:   timeS,
		Value:        val,
		Price:        int64(price),
		Snapshot:     snap,
		PrevSnapshot: prevSnap,
	}
}

func NewPlayerTest() *model.Player {
	addr := "0xFFFbEE9C752A1ca301A3821308F630609a54369A"
	roundAddr := "0x5f63C25ee7f258706803E97d6a866cDe363201c2"
	balance := "850"
	return &model.Player{
		Address:      addr,
		RoundAddress: roundAddr,
		Balance:      balance,
	}
}

func TestNewResp(t *testing.T) {
	lot1 := NewLotTest()
	player := NewPlayerTest()
	lot2 := NewLotTest()

	lot1Resp := &LotResp{
		Name: "lot",
		Data: lot1,
	}

	lot2Resp := &LotResp{
		Name: "lot",
		Data: lot2,
	}

	playerResp := &PlayerResp{
		Name: "player",
		Data: player,
	}

	r := &Resp{}
	r.resp = append(r.resp, lot1Resp)
	r.resp = append(r.resp, lot2Resp)
	r.resp = append(r.resp, playerResp)
	jsonObj, err := json.Marshal(r.resp)
	if err != nil {
		t.Error(err)
	}
	var e interface{}

	if err := json.Unmarshal(jsonObj, &e); err != nil {
		t.Error(err)
	}

	m := e.([]interface{})
	for _, obj := range m {
		newObj, err := obj.(map[string]interface{})
		if !err {
			t.Error(err)
		}
		switch newObj["name"] {
		case "lot":
			lot := &model.Lot{}
			err := mapstructure.Decode(newObj["data"], lot)
			if err != nil {
				t.Error(err)
			}
			fmt.Println(lot)
		case "player":
			player := &model.Player{}
			err := mapstructure.Decode(newObj["data"], player)
			if err != nil {
				t.Error(err)
			}
			fmt.Println(player)
		}
	}
}
