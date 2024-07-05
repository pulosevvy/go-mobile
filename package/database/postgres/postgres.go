package postgres

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"go-mobile/config"
)

type Postgres struct {
	Conn    *pgx.Conn
	Builder squirrel.StatementBuilderType
}

func NewPostgres(cfg *config.PG) (*Postgres, error) {
	const fn = "postgres.New"

	conn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode,
	)

	db, err := pgx.Connect(context.Background(), conn)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	postgres := &Postgres{
		Conn:    db,
		Builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}

	return postgres, nil
}

func (p *Postgres) Close(ctx context.Context) error {
	const fn = "postgres.Close"
	err := p.Conn.Close(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}
	return nil
}
