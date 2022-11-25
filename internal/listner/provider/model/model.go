package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type CreateRoundEvent struct {
	RoundAddr common.Address
}

type EnterRoundEvent struct {
	Sender common.Address
}

type StartRoundEvent struct {
	BalancesSnap *big.Int
	ParamsSnap   *big.Int
}

type CreatedLotEvent struct {
	LotAddr common.Address
}

type NewLotEvent struct {
	LotAddr      common.Address `abi:"_lotAddr"`
	TimeFirst    *big.Int       `abi:"_timeFirst"`
	TimeSecond   *big.Int       `abi:"_timeSecond"`
	Price        *big.Int       `abi:"_price"`
	Value        *big.Int       `abi:"_val"`
	LotSnap      *big.Int       `abi:"_lotSnap"`
	BalancesSnap *big.Int       `abi:"_bsnap"`
}

type BuyLotEvent struct {
	LotAddr      common.Address `abi:"_lotAddr"`
	Sender       common.Address `abi:"_sender"`
	Price        *big.Int       `abi:"_price"`
	LotSnap      *big.Int       `abi:"_lotSnap"`
	BalancesSnap *big.Int       `abi:"_bsnap"`
}

type SendLotEvent struct {
	LotAddr       common.Address `abi:"_lotAddr"`
	ReceiveTokens *big.Int       `abi:"_receiveTokens"`
}

type UpdatePlayerParams struct {
	Owner common.Address
	Nwin  *big.Int
	N     *big.Int
	Spos  *big.Int
	Sneg  *big.Int
}

type ReceiveLotEvent struct {
	LotAddr      common.Address
	Owner        common.Address
	Balance      *big.Int
	ParamsSnap   *big.Int
	BalancesSnap *big.Int
}

var (
	CreateRoundEventHash = crypto.Keccak256Hash([]byte("CreateRoundEvent(address)")).Hex()

	EnterRoundEventHash = crypto.Keccak256Hash([]byte("EnterRoundEvent(address)")).Hex()

	StartRoundEventHash = crypto.Keccak256Hash([]byte("StartRoundEvent(uint256,uint256)")).Hex()

	CreatedLotEventHash = crypto.Keccak256Hash([]byte("CreatedLotEvent(address)")).Hex()

	NewLotEventHash = crypto.Keccak256Hash([]byte("NewLotEvent(address,uint256,uint256,uint256,uint256,uint256,uint256)")).Hex()

	BuyLotEventHash = crypto.Keccak256Hash([]byte("BuyLotEvent(address,address,uint256,uint256,uint256)")).Hex()

	SendLotEventHash = crypto.Keccak256Hash([]byte("SendLotEvent(address,uint256)")).Hex()

	UpdatePlayerParamsHash = crypto.Keccak256Hash([]byte("UpdatePlayerParams(address,uint256,uint256,uint256,uint256)")).Hex()

	ReceiveLotEventHash = crypto.Keccak256Hash([]byte("ReceiveLotEvent(address,address,uint256,uint256,uint256)")).Hex()
)
