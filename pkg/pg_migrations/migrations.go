package pg_migrations

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	pgxv5 "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"log/slog"
)

func Migrate(logger *slog.Logger, pool *pgxpool.Pool) error {
	db := stdlib.OpenDBFromPool(pool)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			logger.Warn("Failed to close DB migration connection", err.Error())
		}
	}(db)

	driver, err := pgxv5.WithInstance(db, &pgxv5.Config{})
	if err != nil {
		return err
	}

	instance, err := migrate.NewWithDatabaseInstance("", "tiny", driver)
	if err != nil {
		return err
	}

	err = instance.Up()
	if err != nil {
		return err
	}

	return nil
}
