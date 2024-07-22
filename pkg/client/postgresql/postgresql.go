package postgresql

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"log"
	"time"
)

const maxAttempts = 5
const delay time.Duration = 2

func NewPgxPool(ctx context.Context, logger *zap.Logger, dsn string) (*pgxpool.Pool, error) {
	pool, err := doWithTries(func() (*pgxpool.Pool, error) {
		pool, err := pgxpool.New(ctx, dsn)
		if err != nil {
			logger.Error("Failed connect pool", zap.Error(err))
			return nil, err
		}
		err = pool.Ping(ctx)
		if err != nil {
			logger.Error("ping: ", zap.Error(err))
			return nil, err
		}
		return pool, nil
	}, maxAttempts, delay)
	if err != nil {
		logger.Error("Error reconnecting:", zap.Error(err))
		return nil, err
	}
	logger.Info("Connect to DB")
	return pool, nil
}

func doWithTries(fn func() (*pgxpool.Pool, error), attempts int, delay time.Duration) (*pgxpool.Pool, error) {
	for attempts > 0 {
		pool, err := fn()
		if err != nil {
			log.Printf("Error connecting to DB: %v", err)
			if attempts == 1 {
				return nil, err
			}
			log.Print("Reconnecting to DB...")
			time.Sleep(delay * time.Second)
			attempts--
			continue
		}
		return pool, nil
	}
	return nil, errors.New("max attempts reached without success")
}
