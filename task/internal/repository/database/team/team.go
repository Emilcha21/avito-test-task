package team

import (
	//"database/sql"
	"log/slog"

	"gorm.io/gorm"
)

type TeamRepoSQL struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewTeamRepoSQL(db *gorm.DB, logger *slog.Logger) *TeamRepoSQL {
	return &TeamRepoSQL{
		db:     db,
		logger: logger,
	}
}
