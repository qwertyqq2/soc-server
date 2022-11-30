package store

import (
	"fmt"
	"log"

	"github.com/qwertyqq2/soc-server/internal/listner/provider/model"
)

type Repository struct {
	store *Store
}

func (r *Repository) CreateRound(e *model.CreateRoundEvent) error {
	query := fmt.Sprintf("INSERT INTO Rounds (address, deposit) VALUES('%s', '%s') ;",
		e.RoundAddr.Hex(),
		e.Deposit.String(),
	)
	_, err := r.store.db.Exec(query)
	return err
}

func (r *Repository) FindRound(roundAddress string) (*model.Round, error) {
	query := fmt.Sprintf("SELECT * FROM Rounds WHERE address='%s';", roundAddress)
	rows, err := r.store.db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	round := &model.Round{}
	for rows.Next() {
		err := rows.Scan(&round.Address, &round.Deposit,
			&round.BalancesSnap, &round.ParamsSnap, &round.Spos, &round.Sneg, &round.Reserve)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return round, nil
}

func (r *Repository) DeleteRound(roundAddress string) error {
	query := fmt.Sprintf("DELETE FROM Rounds WHERE address='%s';", roundAddress)
	_, err := r.store.db.Exec(query)
	return err
}

func (r *Repository) EnterRound(e *model.EnterRoundEvent) error {
	query := fmt.Sprintf("INSERT INTO PendingPlayers (sender, roundAddress) VALUES('%s', '%s') ;",
		e.Sender.Hex(),
		e.RoundAddr.Hex(),
	)
	_, err := r.store.db.Exec(query)
	return err
}

func (r *Repository) NewPlayer(e *model.Player) error {
	query := fmt.Sprintf("INSERT INTO Players (address, roundAddress, balance) VALUES('%s', '%s', '%s') ;",
		e.Address,
		e.RoundAddress,
		e.Balance,
	)
	_, err := r.store.db.Exec(query)
	return err
}

func (r *Repository) FindDepositRound(roundAddress string) (string, error) {
	query := fmt.Sprintf("SELECT * FROM Rounds WHERE address='%s';", roundAddress)
	rows, err := r.store.db.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	round := &model.Round{}
	for rows.Next() {
		err := rows.Scan(&round.Address, &round.Deposit,
			&round.BalancesSnap, &round.ParamsSnap, &round.Spos, &round.Sneg, &round.Reserve)
		if err != nil {
			return "", err
		}
	}
	err = rows.Err()
	if err != nil {
		return "", err
	}
	return round.Deposit, nil
}

func (r *Repository) StartRound(e *model.StartRoundEvent) error {
	// deposit, err := r.FindDepositRound(e.RoundAddr.Hex())
	// if err != nil {
	// 	return err
	// }
	query := fmt.Sprintf("SELECT * FROM PendingPlayers WHERE roundAddress='%s';", e.RoundAddr.Hex())
	rows, err := r.store.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	pendPlayer := &model.PendingPlayer{}
	for rows.Next() {
		err := rows.Scan(&pendPlayer.Sender, &pendPlayer.RoundAddress)
		if err != nil {
			return err
		}
		player := &model.Player{
			Address:      pendPlayer.Sender,
			RoundAddress: pendPlayer.RoundAddress,
			Balance:      "1000",
		}
		if err := r.NewPlayer(player); err != nil {
			return err
		}
		fmt.Println("new player")
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	query = fmt.Sprintf("DELETE FROM PendingPlayers;")
	_, err = r.store.db.Exec(query)
	if err != nil {
		return err
	}
	query = fmt.Sprintf("UPDATE Rounds SET bsnap='%s', psnap='%s' WHERE address='%s'",
		e.BalancesSnap.String(),
		e.ParamsSnap.String(),
		e.RoundAddr.String())
	_, err = r.store.db.Exec(query)
	return err
}

func (r *Repository) CreateLot(e *model.CreatedLotEvent) error {
	query := fmt.Sprintf("INSERT INTO Lots (address, roundAddress) VALUES('%s', '%s') ;",
		e.LotAddr.Hex(),
		e.RoundAddr.Hex(),
	)
	_, err := r.store.db.Exec(query)
	return err
}

func (r *Repository) NewLot(e *model.NewLotEvent) error {
	query := fmt.Sprintf("UPDATE Lots SET owner='%s', timeFirst='%s', timeSecond='%s', value='%s', price='%d', snapshot='%s' WHERE address='%s';",
		e.Owner.Hex(),
		e.TimeFirst.String(),
		e.TimeSecond.String(),
		e.Value.String(),
		e.Price,
		e.LotSnap.String(),
		e.LotAddr.String(),
	)
	_, err := r.store.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	query2 := fmt.Sprintf("UPDATE Rounds SET bsnap='%s' WHERE address='%s';",
		e.BalancesSnap.String(),
		e.RoundAddr.Hex())

	_, err = r.store.db.Exec(query2)
	return err
}

func (r *Repository) BuyLot(e *model.BuyLotEvent) error {
	query := fmt.Sprintf("UPDATE Lots SET owner='%s', price='%d', snapshot='%s' WHERE address='%s';",
		e.Sender.Hex(),
		e.Price,
		e.LotSnap.String(),
		e.LotAddr.String(),
	)
	_, err := r.store.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	query2 := fmt.Sprintf("UPDATE Rounds SET bsnap='%s' WHERE address='%s';",
		e.BalancesSnap.String(),
		e.RoundAddr.Hex())

	_, err = r.store.db.Exec(query2)
	return err
}

func (r *Repository) SendLot(e *model.SendLotEvent) error {
	query := fmt.Sprintf("UPDATE Lots SET receiveTokens='%s' WHERE address='%s';",
		e.ReceiveTokens.String(),
		e.LotAddr.Hex(),
	)
	_, err := r.store.db.Exec(query)
	return err
}

func (r *Repository) UpdatePlayer(e *model.UpdatePlayerParams) error {
	query := fmt.Sprintf("UPDATE Players SET nwin='%d', n='%d', spos='%s', sneg='%s' WHERE address='%s', roundAddress='%s';",
		e.Nwin.Int64(),
		e.N.Int64(),
		e.Spos.String(),
		e.Sneg.String(),
		e.Owner.Hex(),
		e.RoundAddr.Hex(),
	)
	_, err := r.store.db.Exec(query)
	return err
}

func (r *Repository) ReceiveLot(e *model.ReceiveLotEvent) error {
	query := fmt.Sprintf("UPDATE Players SET balance='%s' WHERE address='%s' AND roundAddress='%s';",
		e.Balance.String(),
		e.LotAddr.Hex(),
		e.RoundAddr.Hex(),
	)
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	query = fmt.Sprintf("UPDATE Rounds SET bsnap='%s', psnap='%s' WHERE address='%s';",
		e.BalancesSnap.String(),
		e.ParamsSnap.String(),
		e.RoundAddr.Hex(),
	)
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	query = fmt.Sprintf("UPDATE Lots SET timeF='%s', timeS='%s', value='%s', price='%d', owner='%s', receiveTokens='%s', snapshot='%s' WHERE address='%s';",
		"", "", "", 0, "", "", "", e.RoundAddr.Hex(),
	)
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}
