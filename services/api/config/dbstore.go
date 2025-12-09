package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PgStore struct {
	Config     *Config
	Connection *pgxpool.Pool
}

func NewPgStore(cfg *Config) (store *PgStore, err error) {
	poolConfig, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s sslmode=%s",
			cfg.Dbuser, cfg.Dbpass, cfg.Dbname, cfg.Dbhost, cfg.Dbsslmode,
		),
	)
	if err != nil {
		return store, err
	}
	poolConfig.MaxConns = int32(cfg.PgxPoolMaxconns)
	poolConfig.MinConns = int32(cfg.PgxPoolMinconns)
	poolConfig.MaxConnIdleTime = cfg.PgxPoolMaxconnIdletime

	db, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return store, err
	}

	store = &PgStore{Config: cfg, Connection: db}
	return store, err
}
