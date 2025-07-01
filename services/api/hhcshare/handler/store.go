package handler

import (
	"hhcshare/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

type HandlerStore struct {
	Connection *pgxpool.Pool
	Config     *config.Config
}
