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
	query := fmt.Sprintf("INSERT INTO Rounds (address, deposit) VALUES('%s', '%d') ;",
		e.RoundAddr.Hex(),
		e.Deposit,
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
	query := fmt.Sprintf("INSERT INTO Players (address, roundAddress, balance) VALUES('%s', '%s', '%d') ;",
		e.Address,
		e.RoundAddress,
		e.Balance,
	)
	_, err := r.store.db.Exec(query)
	return err
}

func (r *Repository) FindDepositRound(roundAddress string) (int, error) {
	query := fmt.Sprintf("SELECT * FROM Rounds WHERE address='%s';", roundAddress)
	rows, err := r.store.db.Query(query)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	round := &model.Round{}
	for rows.Next() {
		err := rows.Scan(&round.Address, &round.Deposit,
			&round.BalancesSnap, &round.ParamsSnap, &round.Spos, &round.Sneg, &round.Reserve)
		if err != nil {
			return 0, err
		}
	}

	err = rows.Err()
	if err != nil {
		return 0, err
	}
	return round.Deposit, nil
}

func (r *Repository) StartRound(e *model.StartRoundEvent) error {
	deposit, err := r.FindDepositRound(e.RoundAddr.Hex())
	if err != nil {
		return err
	}
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
			Balance:      deposit,
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
	return nil
}
