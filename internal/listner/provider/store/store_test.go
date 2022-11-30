package store_test

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"testing"
	"time"

	_ "database/sql"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/qwertyqq2/soc-server/internal/listner/provider"
	"github.com/qwertyqq2/soc-server/internal/listner/provider/model"
	"github.com/qwertyqq2/soc-server/internal/listner/provider/store"

	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func GenerateAddress() common.Address {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address
}

func TestCreateRound(t *testing.T) {
	databaseURL := "root:aaa@tcp(127.0.0.1:3306)/SOCDB"

	db, err := provider.NewDB(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	s := store.New(db)

	roundAddr := common.HexToAddress("0x28098151b42770eeAcF328296B082f6Fe8Bc3162")
	deposit := big.NewInt(100)

	event := &model.CreateRoundEvent{
		RoundAddr: roundAddr,
		Deposit:   deposit,
	}

	if err := s.Repository().CreateRound(event); err != nil {
		log.Fatal(err)
	}
	fmt.Println("created")
	fromDB, err := s.Repository().FindRound(roundAddr.Hex())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fromDB.Address, fromDB.Deposit)
}

func TestEnter(t *testing.T) {
	databaseURL := "root:aaa@tcp(127.0.0.1:3306)/SOCDB"
	db, err := provider.NewDB(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	s := store.New(db)

	roundAddr := "0x28098151b42770eeAcF328296B082f6Fe8Bc3162"
	playerAddr1 := "0xFFFbEE9C752A1ca301A3821308F630609a54369A"
	playerAddr2 := "0xA8f38E07FaB155A29a1Fdf42bd0e3433C4513487"
	playerAddr3 := "0x8c2f34394d46A2c4bd09CB43b0c4D2A6e1B7643a"
	playerAddr4 := "0x90ca948986F122B6Ca0E08f383779A8CBE15A55D"

	eventEnter1 := &model.EnterRoundEvent{
		RoundAddr: common.HexToAddress(roundAddr),
		Sender:    common.HexToAddress(playerAddr1),
	}
	eventEnter2 := &model.EnterRoundEvent{
		RoundAddr: common.HexToAddress(roundAddr),
		Sender:    common.HexToAddress(playerAddr2),
	}
	eventEnter3 := &model.EnterRoundEvent{
		RoundAddr: common.HexToAddress(roundAddr),
		Sender:    common.HexToAddress(playerAddr3),
	}
	eventEnter4 := &model.EnterRoundEvent{
		RoundAddr: common.HexToAddress(roundAddr),
		Sender:    common.HexToAddress(playerAddr4),
	}
	err = s.Repository().EnterRound(eventEnter1)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Repository().EnterRound(eventEnter2)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Repository().EnterRound(eventEnter3)
	if err != nil {
		log.Fatal(err)
	}
	err = s.Repository().EnterRound(eventEnter4)
	if err != nil {
		log.Fatal(err)
	}
}

func TestStartRound(t *testing.T) {
	databaseURL := "root:aaa@tcp(127.0.0.1:3306)/SOCDB"
	db, err := provider.NewDB(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	s := store.New(db)
	roundAddr := "0x28098151b42770eeAcF328296B082f6Fe8Bc3162"

	bsnaphash := hex.EncodeToString(solsha3.SoliditySHA3(
		[]string{"address", "uint256"},
		[]interface{}{roundAddr, "10"},
	))

	psnaphash := hex.EncodeToString(solsha3.SoliditySHA3(
		[]string{"address", "uint256"},
		[]interface{}{roundAddr, "20"},
	))

	psnap, err := model.Base64ToInt(psnaphash)
	if err != nil {
		log.Fatal(err)
	}

	bsnap, err := model.Base64ToInt(bsnaphash)
	if err != nil {
		log.Fatal(err)
	}

	startEvent := &model.StartRoundEvent{
		RoundAddr:    common.HexToAddress(roundAddr),
		ParamsSnap:   psnap,
		BalancesSnap: bsnap,
	}
	if err := s.Repository().StartRound(startEvent); err != nil {
		log.Fatal(err)
	}
}

func TestCreateLot(t *testing.T) {
	databaseURL := "root:aaa@tcp(127.0.0.1:3306)/SOCDB"
	db, err := provider.NewDB(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	s := store.New(db)
	roundAddr := common.HexToAddress("0x28098151b42770eeAcF328296B082f6Fe8Bc3162")
	lotAddr := common.HexToAddress("0xcd234a471b72ba2f1ccf0a70fcaba648a5eecd8d")

	createLotEvent := &model.CreatedLotEvent{
		LotAddr:   lotAddr,
		RoundAddr: roundAddr,
	}
	if err := s.Repository().CreateLot(createLotEvent); err != nil {
		log.Fatal(err)
	}
}

func TestNewLot(t *testing.T) {
	databaseURL := "root:aaa@tcp(127.0.0.1:3306)/SOCDB"
	db, err := provider.NewDB(databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	s := store.New(db)
	roundAddr := common.HexToAddress("0x28098151b42770eeAcF328296B082f6Fe8Bc3162")
	lotAddr := common.HexToAddress("0xcd234a471b72ba2f1ccf0a70fcaba648a5eecd8d")

	lotsnap, err := model.Base64ToInt(hex.EncodeToString(solsha3.SoliditySHA3(
		[]string{"address", "uint256"},
		[]interface{}{"0xA8f38E07FaB155A29a1Fdf42bd0e3433C4513487", "200"},
	)))
	if err != nil {
		log.Fatal(err)
	}
	bsnap, err := model.Base64ToInt(hex.EncodeToString(solsha3.SoliditySHA3(
		[]string{"address", "uint256"},
		[]interface{}{roundAddr, "10"},
	)))
	if err != nil {
		log.Fatal(err)
	}

	curtime := time.Now().UTC().UnixNano() + 30

	newLotEvent := &model.NewLotEvent{
		RoundAddr:    roundAddr,
		LotAddr:      lotAddr,
		Owner:        common.HexToAddress("0xA8f38E07FaB155A29a1Fdf42bd0e3433C4513487"),
		TimeFirst:    big.NewInt(curtime + 30),
		TimeSecond:   big.NewInt(curtime + 120),
		Price:        big.NewInt(200),
		Value:        big.NewInt(10000000000),
		LotSnap:      lotsnap,
		BalancesSnap: bsnap,
	}

	if err := s.Repository().NewLot(newLotEvent); err != nil {
		log.Fatal(err)
	}

}
