package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

// Interface để dễ mock trong test
type AddNumbersStore interface {
	AddNumbers(ctx context.Context, a, b int) (int, error)
}

func NewPostgres(user, password, host, dbname string, port int) (*Postgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user, password, host, port, dbname)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	config.ConnConfig.RuntimeParams["search_path"] = "public"

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create db pool: %w", err)
	}

	fmt.Println("✅ Connected to PostgreSQL:", dbname, "/ schema=public")
	return &Postgres{Pool: pool}, nil
}

func (p *Postgres) Close() {
	p.Pool.Close()
}

func (p *Postgres) AddNumbers(ctx context.Context, a, b int) (int, error) {
	var result int
	err := p.Pool.QueryRow(ctx, "SELECT add_numbers_hard($1, $2)", a, b).Scan(&result)
	if err != nil {
		return 0, fmt.Errorf("query error: %w", err)
	}
	return result, nil
}
