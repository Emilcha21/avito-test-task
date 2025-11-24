package pullrequest

import (
	"avito-test/internal/models"
	"time"

	"gorm.io/gorm"
)

func (r *PullRequestRepoSQL) SetMergeStatus(prId string, mergedAt time.Time) error {
	result := r.db.Model(&models.PullRequest{}).
		Where("pull_request_id = ?", prId).
		Updates(map[string]interface{}{
			"status":    "MERGED",
			"merged_at": mergedAt,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
