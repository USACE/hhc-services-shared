package handlers

import (
	"context"
	"fmt"
	"hhcshare/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	Connection *pgxpool.Pool
	Config     *config.Config
}

func (s *Store) DbConnection() error {
	poolConfig, err := pgxpool.ParseConfig(
		fmt.Sprintf(
			"user=%s password=%s dbname=%s host=%s sslmode=%s",
			s.Config.DbUser,
			s.Config.DbPass,
			s.Config.DbName,
			s.Config.DbHost,
			s.Config.DbSslMode,
		),
	)
	if err != nil {
		return err
	}

	poolConfig.MaxConns = int32(s.Config.DbPoolMaxConns)
	poolConfig.MinConns = int32(s.Config.DbPoolMinConns)
	poolConfig.MaxConnIdleTime = s.Config.DbPoolMaxConnIdleTime

	connection, err := pgxpool.ConnectConfig(context.TODO(), poolConfig)

	s.Connection = connection

	return err
}

// func (s *Store) SetStoreConfig(cfg *config.Config) {
// 	s.Config = cfg
// }
