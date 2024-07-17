package configs

import (
	"Testovoe/pkg/client/postgresql"
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

type dbParam struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Load(path string, logger *zap.Logger) (*pgxpool.Pool, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	dsn := newDSN()
	pool, err := postgresql.NewPgxPool(context.Background(), logger, dsn)
	if err != nil {
		logger.Fatal("Failed connect pool:", zap.Error(err))
		return nil, err
	}
	m, err := migrate.New("file://app/migrations", dsn+"?sslmode=disable")
	if err != nil {
		logger.Fatal("Error migrate: ", zap.Error(err))
		return nil, err
	}
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatal("Error up: ", zap.Error(err))
		return nil, err
	}
	logger.Info("Migrations UP")

	return pool, nil
}

func newDSN() string {
	dbPar := loadParam()
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		dbPar.User, dbPar.Password, dbPar.Host, dbPar.Port, dbPar.DBName)
}

func loadParam() *dbParam {
	return &dbParam{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}
