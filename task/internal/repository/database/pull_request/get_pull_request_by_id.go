package pullrequest

import (
	"avito-test/internal/models"
	"errors"

	"gorm.io/gorm"
)

func (r *PullRequestRepoSQL) GetPullRequestById(id string) (*models.PullRequest, error) {
	var model models.PullRequest

	result := r.db.
		Select("pull_request_id, pull_request_name, author_id, status, merged_at").
		Where("pull_request_id = ?", id).
		First(&model)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, result.Error
	}

	return &model, nil
}
