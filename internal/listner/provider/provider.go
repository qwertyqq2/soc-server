package provider

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/qwertyqq2/soc-server/apigroup"
	"github.com/qwertyqq2/soc-server/internal/listner/provider/model"
)

type Provider struct {
	client          *ethclient.Client
	contractAddress common.Address
	api             *apigroup.Apigroup
}

func (p *Provider) SetUp() {
	conf := GetConfig()

	pClient, err := ethclient.Dial(conf.ParamsProvider.ProviderURL)
	if err != nil {
		log.Fatal(err)
	}
	p.client = pClient
	if err != nil {
		log.Fatal(err)
	}
	p.contractAddress = common.HexToAddress(conf.ParamsProvider.GroupContractAddress)
	pContractClient, err := apigroup.NewApigroup(p.contractAddress, p.client)
	if err != nil {
		log.Fatal(err)
	}
	p.api = pContractClient
	log.Println("Provider created: ", p.contractAddress.Hex())
}

func (p *Provider) ListenEvent() {
	groupAbi, err := abi.JSON(strings.NewReader(string(apigroup.ApigroupABI)))
	if err != nil {
		log.Fatal(err)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{p.contractAddress},
	}
	ctx := context.Background()

	var eventCh = make(chan types.Log)

	sub, err := p.client.SubscribeFilterLogs(ctx, query, eventCh)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case err := <-sub.Err():
			log.Fatal(err)
		case logEvent := <-eventCh:
			switch logEvent.Topics[0].Hex() {
			case model.CreateRoundEventHash:
				createRoundEvent := &model.CreateRoundEvent{}
				err := groupAbi.UnpackIntoInterface(createRoundEvent, "CreateRoundEvent", logEvent.Data)
				if err != nil {
					log.Println(err)
				}
				log.Println("event-create round: ", createRoundEvent.RoundAddr)

			case model.StartRoundEventHash:
				startRoundEvent := &model.StartRoundEvent{}
				err := groupAbi.UnpackIntoInterface(startRoundEvent, "StartRoundEvent", logEvent.Data)
				if err != nil {
					log.Println(err)
				}
				log.Println("event-start round: ", startRoundEvent.BalancesSnap, startRoundEvent.ParamsSnap)

			case model.EnterRoundEventHash:
				enterRoundEvent := &model.EnterRoundEvent{}
				err := groupAbi.UnpackIntoInterface(enterRoundEvent, "EnterRoundEvent", logEvent.Data)
				if err != nil {
					log.Println(err)
				}
				log.Println("event-enter round: ", enterRoundEvent.Sender)

			case model.NewLotEventHash:
				newLotEvent := &model.NewLotEvent{}
				err := groupAbi.UnpackIntoInterface(newLotEvent, "NewLotEvent", logEvent.Data)
				if err != nil {
					log.Println(err)
				}
				log.Println("event: new lot ", newLotEvent)

			case model.BuyLotEventHash:
				buyLotEvent := &model.BuyLotEvent{}
				err := groupAbi.UnpackIntoInterface(buyLotEvent, "BuyLotEvent", logEvent.Data)
				if err != nil {
					log.Println(err)
				}
				log.Println("event: buy lot ", buyLotEvent)

			case model.SendLotEventHash:
				sendLotEvent := &model.SendLotEvent{}
				err := groupAbi.UnpackIntoInterface(sendLotEvent, "SendLotEvent", logEvent.Data)
				if err != nil {
					fmt.Println(err)
				}
				log.Println("event: send lot ", sendLotEvent)
			}

		}
	}
}
