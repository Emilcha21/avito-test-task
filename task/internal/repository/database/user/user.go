package user

import (
	"log/slog"

	"gorm.io/gorm"
)

type UserRepoSQL struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewUserRepoSQL(db *gorm.DB, logger *slog.Logger) *UserRepoSQL {
	return &UserRepoSQL{
		db:     db,
		logger: logger,
	}
}
