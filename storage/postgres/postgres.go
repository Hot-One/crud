package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"user/config"
	"user/storage"
)

type store struct {
	db   *sql.DB
	user *userRepo
}

func NewConnectionPostgres(cfg config.Config) (storage.StorageI, error) {
	connectionStr := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	)

	sqlDb, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}

	return &store{
		db:   sqlDb,
		user: NewUserRepo(sqlDb),
	}, nil
}

func (s *store) Close() {
	s.db.Close()
}

func (s *store) User() storage.UserRepoI {
	s.user = NewUserRepo(s.db)
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}
