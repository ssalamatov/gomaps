package server

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/ssalamatov/gomaps/internal/config"
	"github.com/ssalamatov/gomaps/pkg/client/postgresql"
)

type Server struct {
	pool *pgxpool.Pool
	ctx  context.Context
}

func NewServer(config *config.Config) (*Server, error) {
	ctx := context.Background()

	pool, err := postgresql.NewPgClient(ctx, config)
	if err != nil {
		return nil, err
	}
	return &Server{pool: pool, ctx: ctx}, nil
}
