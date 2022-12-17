package model

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/mitchellh/mapstructure"
)

func NewLotTest() *LotWrapper {
	addr := "0xfB3Ee894f27d8d6798f0034d7763b7ad46C31B92"
	roundAddr := "0x5f63C25ee7f258706803E97d6a866cDe363201c2"
	owner := "0xFFFbEE9C752A1ca301A3821308F630609a54369A"
	timeF := "123345534234234"
	timeS := "133345534234234"
	val := "20"
	price := 60
	snap := "15468303776996484186685289508981412665698293446761833792330076793367195714725"
	prevSnap := "57003951752417535200923101159811040833146496380398724999470920633816778075019"
	lot := &Lot{
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
	lotWrapper := NewLotWrapper(lot)
	return lotWrapper
}

func NewPlayerTest() *PlayerWrapper {
	addr := "0xFFFbEE9C752A1ca301A3821308F630609a54369A"
	roundAddr := "0x5f63C25ee7f258706803E97d6a866cDe363201c2"
	balance := "850"
	player := &Player{
		Address:      addr,
		RoundAddress: roundAddr,
		Balance:      balance,
	}
	playerWrapper := NewPlayerWrapper(player)
	return playerWrapper
}

func NewRespTest(data ...interface{}) *Resp {
	return &Resp{
		resp: data,
	}
}

func UnmarshalResp(jsonResp []byte) {
	var e interface{}
	if err := json.Unmarshal(jsonResp, &e); err != nil {
		log.Fatal(err)
	}
	m, ok := e.([]interface{})
	if !ok {
		log.Println("not ok")
	}
	for _, obj := range m {
		newObj, ok := obj.(map[string]interface{})
		if !ok {
			log.Fatal(ok)
		}
		switch newObj["name"] {
		case "lot":
			lot := &Lot{}
			err := mapstructure.Decode(newObj["data"], lot)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("LOT", lot)
		case "player":
			player := &Player{}
			err := mapstructure.Decode(newObj["data"], player)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("PLAYER", player)
		}
	}
}

func TestRespInput(t *testing.T) {
	lot := NewLotTest()
	lot2 := NewLotTest()
	//lots := []*Lot{lot, lot2}
	player := NewPlayerTest()
	resp := NewRespTest(lot, player, lot2)
	jsonResp, _ := resp.Marshal()
	UnmarshalResp(jsonResp)
}
