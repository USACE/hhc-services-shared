package handler

import (
	"github.com/hhc-services-shared/services/api/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

type HandlerStore struct {
	Connection *pgxpool.Pool
	Config     *config.Config
}
