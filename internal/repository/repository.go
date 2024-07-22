package repository

import (
	"Testovoe/internal/domen"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name=Repository
type Repository interface {
	Create(context.Context, *domen.Response) error
}

type Repo struct {
	pg     *pgxpool.Pool
	logger *zap.Logger
	tp     trace.Tracer
}

func NewRepo(pgx *pgxpool.Pool, logger *zap.Logger, tp trace.Tracer) *Repo {
	return &Repo{pgx, logger, tp}
}

func (r *Repo) Create(ctx context.Context, respDTO *domen.ResponseDTO) error {
	_, span := r.tp.Start(ctx, "repo create ")
	defer span.End()
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
