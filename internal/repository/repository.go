package repository

import (
	"Testovoe/internal/domen"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Repository interface {
	Create(context.Context, *domen.Response) error
}

type Repo struct {
	pg     *pgxpool.Pool
	logger *zap.Logger
}

func NewRepo(pgx *pgxpool.Pool, logger *zap.Logger) *Repo {
	return &Repo{pgx, logger}
}

func (r *Repo) Create(ctx context.Context, respDTO *domen.ResponseDTO) error {

	_, err := r.pg.Exec(ctx, `
			INSERT INTO asks (timestamp, price, volume, amount, factor, type)
			VALUES ($1, $2, $3, $4, $5, $6)
			`, respDTO.Timestamp, respDTO.Asks.Price, respDTO.Asks.Volume,
		respDTO.Asks.Amount, respDTO.Asks.Factor, respDTO.Asks.Type)
	if err != nil {
		return err
	}

	_, err = r.pg.Exec(ctx, `
			INSERT INTO bids (timestamp, price, volume, amount, factor, type)
			VALUES ($1, $2, $3, $4, $5, $6)
			`, respDTO.Timestamp, respDTO.Bids.Price, respDTO.Bids.Volume,
		respDTO.Bids.Amount, respDTO.Bids.Factor, respDTO.Bids.Type)
	if err != nil {
		return err
	}

	return nil
}
