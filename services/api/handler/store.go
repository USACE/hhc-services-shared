package handler

import (
	"github.com/hhc-services-shared/services/api/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

type HandlerStore struct {
	Connection *pgxpool.Pool
	Config     *config.Config
}
