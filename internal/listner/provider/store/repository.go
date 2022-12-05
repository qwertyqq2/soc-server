package store

import (
	"fmt"
	"log"
	"math/big"

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
	query := fmt.Sprintf("SELECT * FROM PendingPlayers WHERE roundAddress='%s';", e.RoundAddr.Hex())
	rows, err := r.store.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	pendPlayer := &model.PendingPlayer{}
	deposit, err := r.FindDepositRound(e.RoundAddr.Hex())
	if err != nil {
		return err
	}
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

func (r *Repository) FindPlayerBalance(playerAddr string) (*big.Int, error) {
	query := fmt.Sprintf("SELECT * FROM Players WHERE address='%s';", playerAddr)
	rows, err := r.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	player := &model.Player{}
	for rows.Next() {
		err := rows.Scan(&player.Address, &player.RoundAddress, &player.Balance, &player.Nwin,
			&player.N, &player.Spos, &player.Sneg)
		if err != nil {
			return nil, err
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	balance := new(big.Int)
	balance.SetString(player.Balance, 10)
	return balance, nil
}

func (r *Repository) NewLot(e *model.NewLotEvent) error {
	prevBalance, err := r.FindPlayerBalance(e.Owner.Hex())
	if err != nil {
		return err
	}
	query := fmt.Sprintf("UPDATE Lots SET owner='%s', timeFirst='%s', timeSecond='%s', value='%s', price='%d', snapshot='%s' WHERE address='%s';",
		e.Owner.Hex(),
		e.TimeFirst.String(),
		e.TimeSecond.String(),
		e.Value.String(),
		e.Price,
		e.LotSnap.String(),
		e.LotAddr.String(),
	)
	_, err = r.store.db.Exec(query)
	if err != nil {
		return err
	}
	newBalance := big.NewInt(0)
	newBalance.Sub(prevBalance, e.Price)
	query2 := fmt.Sprintf("UPDATE Rounds SET bsnap='%s' WHERE address='%s';",
		e.BalancesSnap.String(),
		e.RoundAddr.Hex())
	_, err = r.store.db.Exec(query2)
	if err != nil {
		return err
	}
	query3 := fmt.Sprintf("UPDATE Players SET balance='%s' WHERE address='%s';",
		newBalance.String(),
		e.Owner.Hex())
	_, err = r.store.db.Exec(query3)
	return err
}

func (r *Repository) BuyLot(e *model.BuyLotEvent) error {
	prevBalance, err := r.FindPlayerBalance(e.Sender.Hex())
	if err != nil {
		return err
	}
	query := fmt.Sprintf("UPDATE Lots SET owner='%s', price='%d', snapshot='%s' WHERE address='%s';",
		e.Sender.Hex(),
		e.Price,
		e.LotSnap.String(),
		e.LotAddr.String(),
	)
	_, err = r.store.db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	newBalance := big.NewInt(0)
	newBalance.Sub(prevBalance, e.Price)

	query2 := fmt.Sprintf("UPDATE Rounds SET bsnap='%s' WHERE address='%s';",
		e.BalancesSnap.String(),
		e.RoundAddr.Hex())
	_, err = r.store.db.Exec(query2)
	if err != nil {
		return err
	}
	query3 := fmt.Sprintf("UPDATE Players SET balance='%s' WHERE address='%s';",
		newBalance.String(),
		e.Sender.Hex())
	_, err = r.store.db.Exec(query3)
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
	query := fmt.Sprintf("UPDATE Players SET nwin='%d', n='%d', spos='%s', sneg='%s' WHERE address='%s' AND roundAddress='%s';",
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

func (r *Repository) FindRoundSpos(roundAddr string) (*big.Int, error) {
	query := fmt.Sprintf("SELECT * FROM Rounds WHERE address='%s';", roundAddr)
	rows, err := r.store.db.Query(query)
	if err != nil {
		return nil, err
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
	Spos := new(big.Int)
	Spos.SetString(round.Spos, 10)
	return Spos, nil
}

func (r *Repository) FindRoundSneg(roundAddr string) (*big.Int, error) {
	query := fmt.Sprintf("SELECT * FROM Rounds WHERE address='%s';", roundAddr)
	rows, err := r.store.db.Query(query)
	if err != nil {
		return nil, err
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
	Sneg := new(big.Int)
	Sneg.SetString(round.Sneg, 10)
	return Sneg, nil
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
	Spos, err := r.FindRoundSpos(e.RoundAddr.Hex())
	if err != nil {
		log.Fatal(err)
	}
	Sneg, err := r.FindRoundSneg(e.RoundAddr.Hex())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Spos.String())

	newSpos := big.NewInt(0)
	newSneg := big.NewInt(0)
	newSpos.Add(Spos, e.SposDelta)
	newSneg.Add(Sneg, e.SnegDelta)
	fmt.Println(newSpos.String())
	query = fmt.Sprintf("UPDATE Rounds SET bsnap='%s', psnap='%s', Spos='%s', Sneg='%s' WHERE address='%s';",
		e.BalancesSnap.String(),
		e.ParamsSnap.String(),
		newSpos.String(),
		newSneg.String(),
		e.RoundAddr.Hex(),
	)
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	query = fmt.Sprintf("UPDATE Lots SET timeFirst='%s', timeSecond='%s', value='%s', price='%d', owner='%s', receiveTokens='%s', snapshot='%s' WHERE address='%s';",
		"", "", "", 0, "", "", "", e.RoundAddr.Hex(),
	)
	if _, err := r.store.db.Exec(query); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Clear() error {
	query := fmt.Sprintf("delete from Players;")
	_, err := r.store.db.Exec(query)
	query = fmt.Sprintf("delete from Rounds;")
	_, err = r.store.db.Exec(query)
	query = fmt.Sprintf("delete from Lots;")
	_, err = r.store.db.Exec(query)
	query = fmt.Sprintf("delete from PendingPlayers;")
	_, err = r.store.db.Exec(query)
	return err
}

func (r *Repository) AllPlayers() ([]*model.Player, error) {
	query := fmt.Sprintf("SELECT * FROM Players;")
	rows, err := r.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	players := make([]*model.Player, 0)
	for rows.Next() {
		player := &model.Player{}
		err := rows.Scan(&player.Address, &player.RoundAddress, &player.Balance, &player.Nwin,
			&player.N, &player.Spos, &player.Sneg)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return players, nil
}

func (r *Repository) AllRounds() ([]*model.Round, error) {
	query := fmt.Sprintf("SELECT * FROM Rounds;")
	rows, err := r.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	rounds := make([]*model.Round, 0)
	for rows.Next() {
		round := &model.Round{}
		err := rows.Scan(&round.Address, &round.Deposit,
			&round.BalancesSnap, &round.ParamsSnap, &round.Spos, &round.Sneg, &round.Reserve)
		if err != nil {
			return nil, err
		}
		rounds = append(rounds, round)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return rounds, nil
}

func (r *Repository) AllLots() ([]*model.Lot, error) {
	query := fmt.Sprintf("SELECT * FROM Lots;")
	rows, err := r.store.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	lots := make([]*model.Lot, 0)
	for rows.Next() {
		lot := &model.Lot{}
		err := rows.Scan(&lot.Address, &lot.RoundAddress,
			&lot.TimeFirst, &lot.TimeSecond, &lot.Value, &lot.Price,
			&lot.Owner, &lot.ReceiveTokens, &lot.Snapshot)
		if err != nil {
			return nil, err
		}
		lots = append(lots, lot)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return lots, nil
}
