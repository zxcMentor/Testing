package repository

import (
	"Testovoe/internal/domen"
	"database/sql"
	"fmt"
)

type Repository interface {
	Create(respMarket *domen.Response) error
}

type Repo struct {
	Db sql.DB
}

func (r *Repo) Create(respMarket *domen.Response) error {
	tx, err := r.Db.Begin()
	if err != nil {
		fmt.Println("err begin tx")
	}
	for _, bid := range respMarket.Bids {
		_, err = tx.Exec("INSERT INTO bids (timestamp, price, volume, amount, factor, type) VALUES ($1, $2, $3, $4, $5, $6)",
			respMarket.Timestamp, bid.Price, bid.Volume,
			bid.Amount, bid.Factor, bid.Type)

		if err != nil {
			tx.Rollback()
			return err
		}
	}
	for _, ask := range respMarket.Ask {
		_, err = tx.Exec("INSERT INTO asks (timestamp, price, volume, amount, factor, type) VALUES ($1, $2, $3, $4, $5, $6)",
			respMarket.Timestamp, ask.Price, ask.Volume,
			ask.Amount, ask.Factor, ask.Type)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return nil
}
