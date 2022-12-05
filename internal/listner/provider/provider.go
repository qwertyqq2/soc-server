package provider

import (
	"context"
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/qwertyqq2/soc-server/apigroup"
	"github.com/qwertyqq2/soc-server/internal/configs"
	"github.com/qwertyqq2/soc-server/internal/listner/provider/model"
	"github.com/qwertyqq2/soc-server/internal/listner/provider/store"
)

type Provider struct {
	client          *ethclient.Client
	contractAddress common.Address
	api             *apigroup.Apigroup
	Store           *store.Store
}

func NewDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (p *Provider) SetUp() error {
	conf := configs.GetConfig()
	db, err := NewDB(conf.DB.DbUrl)
	if err != nil {
		return err
	}

	p.Store = store.New(db)

	pClient, err := ethclient.Dial(conf.ParamsProvider.ProviderURL)
	if err != nil {
		return err
	}
	p.client = pClient
	if err != nil {
		return err
	}
	p.contractAddress = common.HexToAddress(conf.ParamsProvider.GroupContractAddress)
	pContractClient, err := apigroup.NewApigroup(p.contractAddress, p.client)
	if err != nil {
		return err
	}
	p.api = pContractClient
	log.Println("Provider created: ", p.contractAddress.Hex())
	return nil
}

func (p *Provider) ListenEvent() error {
	defer p.Store.Close()
	if err := p.Store.Repository().Clear(); err != nil {
		log.Fatal(err)
	}
	groupAbi, err := abi.JSON(strings.NewReader(string(apigroup.ApigroupABI)))
	if err != nil {
		return err
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{p.contractAddress},
	}
	ctx := context.Background()
	var eventCh = make(chan types.Log)
	sub, err := p.client.SubscribeFilterLogs(ctx, query, eventCh)
	if err != nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-sub.Err():
			return err
		case logEvent := <-eventCh:
			switch logEvent.Topics[0].Hex() {
			case model.CreateRoundEventHash:
				createRoundEvent := &model.CreateRoundEvent{}
				err := groupAbi.UnpackIntoInterface(createRoundEvent, "CreateRoundEvent", logEvent.Data)
				if err != nil {
					return err
				}
				if err := p.Store.Repository().CreateRound(createRoundEvent); err != nil {
					return err
				}
				log.Println("event-create round: ", createRoundEvent)

			case model.EnterRoundEventHash:
				enterRoundEvent := &model.EnterRoundEvent{}
				err := groupAbi.UnpackIntoInterface(enterRoundEvent, "EnterRoundEvent", logEvent.Data)
				if err != nil {
					return err
				}
				if err := p.Store.Repository().EnterRound(enterRoundEvent); err != nil {
					return err
				}
				log.Println("event-enter round: ", enterRoundEvent)

			case model.StartRoundEventHash:
				startRoundEvent := &model.StartRoundEvent{}
				err := groupAbi.UnpackIntoInterface(startRoundEvent, "StartRoundEvent", logEvent.Data)
				if err != nil {
					return err
				}
				if err := p.Store.Repository().StartRound(startRoundEvent); err != nil {
					return err
				}
				log.Println("event-start round: ", startRoundEvent)

			case model.CreatedLotEventHash:
				createLotEvent := &model.CreatedLotEvent{}
				err := groupAbi.UnpackIntoInterface(createLotEvent, "CreatedLotEvent", logEvent.Data)
				if err != nil {
					return err
				}
				if err := p.Store.Repository().CreateLot(createLotEvent); err != nil {
					return err
				}
				log.Println("event-create lot: ", createLotEvent)

			case model.NewLotEventHash:
				newLotEvent := &model.NewLotEvent{}
				err := groupAbi.UnpackIntoInterface(newLotEvent, "NewLotEvent", logEvent.Data)
				if err != nil {
					return err
				}
				if err := p.Store.Repository().NewLot(newLotEvent); err != nil {
					return err
				}
				log.Println("event-new lot ", newLotEvent)

			case model.BuyLotEventHash:
				buyLotEvent := &model.BuyLotEvent{}
				err := groupAbi.UnpackIntoInterface(buyLotEvent, "BuyLotEvent", logEvent.Data)
				if err != nil {
					return err
				}
				if err := p.Store.Repository().BuyLot(buyLotEvent); err != nil {
					return err
				}
				log.Println("event-buy lot ", buyLotEvent)

			case model.SendLotEventHash:
				sendLotEvent := &model.SendLotEvent{}
				err := groupAbi.UnpackIntoInterface(sendLotEvent, "SendLotEvent", logEvent.Data)
				if err != nil {
					return err
				}
				if err := p.Store.Repository().SendLot(sendLotEvent); err != nil {
					return err
				}
				log.Println("event-send lot ", sendLotEvent)

			case model.UpdatePlayerParamsHash:
				updateParamsEvent := &model.UpdatePlayerParams{}
				err := groupAbi.UnpackIntoInterface(updateParamsEvent, "UpdatePlayerParams", logEvent.Data)
				if err != nil {
					return err
				}
				if err := p.Store.Repository().UpdatePlayer(updateParamsEvent); err != nil {
					return err
				}
				log.Println("event-update params ", updateParamsEvent)

			case model.ReceiveLotEventHash:
				receiveLotEvent := &model.ReceiveLotEvent{}
				err := groupAbi.UnpackIntoInterface(receiveLotEvent, "ReceiveLotEvent", logEvent.Data)
				if err != nil {
					return err
				}
				if err := p.Store.Repository().ReceiveLot(receiveLotEvent); err != nil {
					return err
				}
				log.Println("event-receive lot ", receiveLotEvent)
			}

		}
	}
}
