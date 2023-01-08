package migrate

import (
	"database/sql"
	"embed"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/pressly/goose/v3"
	"github.com/sabirovruslan/otus_golang/hw12_13_14_15_calendar/internal/config"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

type Migrate interface {
	Run() error
}

type PgMigrate struct {
	db      *sql.DB
	dialect string
	dir     string
}

func NewPgMigrate(c *config.DatabaseConf) (Migrate, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("migrate create db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("migrate ping db: %w", err)
	}
	return &PgMigrate{db: db, dialect: c.Dialect, dir: c.DirMigrate}, nil
}

func (m *PgMigrate) Run() error {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect(m.dialect); err != nil {
		return fmt.Errorf("migrate dialect: %w", err)
	}

	if err := goose.Up(m.db, m.dir); err != nil {
		return fmt.Errorf("migrate up: %w", err)
	}
	return nil
}
