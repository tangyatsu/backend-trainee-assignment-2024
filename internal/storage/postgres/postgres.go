package postgres

import (
	"backend-trainee-assignment-2024/internal/config"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

func InitDB(cfg *config.Config) (*sqlx.DB, error) {
	const op = "storage.postgres.InitDB"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DBName, cfg.Postgres.SSLMode)

	db, err := sqlx.Connect("pgx", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Not optimal, slow on filtering by tag, separate it to more relations.
	stmt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "banner"
		(
			id SERIAL  PRIMARY KEY,
			feature_id INT NOT NULL,
			tag_ids INT[] NOT NULL,
			is_active BOOL NOT NULL DEFAULT TRUE,
			content text NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);
		CREATE INDEX IF NOT EXISTS banner_idx ON banner (id);
	`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return db, nil
}
