package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type CreateRoundEvent struct {
	RoundAddr common.Address `abi:"_roundAddress"`
	Deposit   *big.Int       `abi:"_deposit"`
}

type EnterRoundEvent struct {
	RoundAddr common.Address `abi:"_roundAddress"`
	Sender    common.Address `abi:"_sender"`
}

type StartRoundEvent struct {
	RoundAddr    common.Address `abi:"_roundAddress"`
	BalancesSnap *big.Int       `abi:"_bsnap"`
	ParamsSnap   *big.Int       `abi:"_psnap"`
}

type CreatedLotEvent struct {
	RoundAddr common.Address `abi:"_roundAddress"`
	LotAddr   common.Address `abi:"_lotAddr"`
}

type NewLotEvent struct {
	RoundAddr    common.Address `abi:"_roundAddress"`
	LotAddr      common.Address `abi:"_lotAddr"`
	Owner        common.Address `abi:"_owner"`
	TimeFirst    *big.Int       `abi:"_timeFirst"`
	TimeSecond   *big.Int       `abi:"_timeSecond"`
	Price        *big.Int       `abi:"_price"`
	Value        *big.Int       `abi:"_val"`
	LotSnap      *big.Int       `abi:"_lotSnap"`
	BalancesSnap *big.Int       `abi:"_bsnap"`
}

type BuyLotEvent struct {
	RoundAddr    common.Address `abi:"_roundAddress"`
	LotAddr      common.Address `abi:"_lotAddr"`
	Sender       common.Address `abi:"_sender"`
	Price        *big.Int       `abi:"_price"`
	LotSnap      *big.Int       `abi:"_lotSnap"`
	BalancesSnap *big.Int       `abi:"_bsnap"`
}

type SendLotEvent struct {
	RoundAddr     common.Address `abi:"_roundAddress"`
	LotAddr       common.Address `abi:"_lotAddr"`
	ReceiveTokens *big.Int       `abi:"_receiveTokens"`
}

type UpdatePlayerParams struct {
	RoundAddr common.Address `abi:"_roundAddress"`
	Owner     common.Address `abi:"_owner"`
	Nwin      *big.Int       `abi:"_nwin"`
	N         *big.Int       `abi:"_n"`
	Spos      *big.Int       `abi:"_spos"`
	Sneg      *big.Int       `abi:"_sneg"`
}

type ReceiveLotEvent struct {
	RoundAddr    common.Address `abi:"_roundAddress"`
	LotAddr      common.Address `abi:"_lotAddr"`
	Owner        common.Address `abi:"_owner"`
	Balance      *big.Int       `abi:"_balance"`
	SposDelta    *big.Int       `abi:"_SposDelta"`
	SnegDelta    *big.Int       `abi:"_SnegDelta"`
	ParamsSnap   *big.Int       `abi:"_psnap"`
	BalancesSnap *big.Int       `abi:"_bsnap"`
}

var (
	CreateRoundEventHash = crypto.Keccak256Hash([]byte("CreateRoundEvent(address,uint256)")).Hex()

	EnterRoundEventHash = crypto.Keccak256Hash([]byte("EnterRoundEvent(address,address)")).Hex()

	StartRoundEventHash = crypto.Keccak256Hash([]byte("StartRoundEvent(address,uint256,uint256)")).Hex()

	CreatedLotEventHash = crypto.Keccak256Hash([]byte("CreatedLotEvent(address,address)")).Hex()

	NewLotEventHash = crypto.Keccak256Hash([]byte("NewLotEvent(address,address,address,uint256,uint256,uint256,uint256,uint256,uint256)")).Hex()

	BuyLotEventHash = crypto.Keccak256Hash([]byte("BuyLotEvent(address,address,address,uint256,uint256,uint256)")).Hex()

	SendLotEventHash = crypto.Keccak256Hash([]byte("SendLotEvent(address,address,uint256)")).Hex()

	UpdatePlayerParamsHash = crypto.Keccak256Hash([]byte("UpdatePlayerParams(address,address,uint256,uint256,uint256,uint256)")).Hex()

	ReceiveLotEventHash = crypto.Keccak256Hash([]byte("ReceiveLotEvent(address,address,address,uint256,uint256,uint256,uint256,uint256)")).Hex()
)
