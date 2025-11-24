package pullrequest

import (
	"log/slog"

	"gorm.io/gorm"
)

type PullRequestRepoSQL struct {
	db     *gorm.DB
	logger *slog.Logger
}

func NewPullRequestRepoSQL(db *gorm.DB, logger *slog.Logger) *PullRequestRepoSQL {
	return &PullRequestRepoSQL{
		db:     db,
		logger: logger,
	}
}
