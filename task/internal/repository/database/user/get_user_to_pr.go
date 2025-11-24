package user

import (
	"avito-test/consts"
	"avito-test/internal/models"

	"gorm.io/gorm"
)

func (r *UserRepoSQL) GetUserToPr(teamName string, authorId string, prId string, oldId string) (string, error) {
	var userId string

	subQuery := r.db.Table("pr_users").
		Select("user_id").
		Where("pull_request_id = ?", prId)

	result := r.db.Model(&models.User{}).
		Select("user_id").
		Where("team_name = ? AND is_active = TRUE", teamName).
		Where("user_id != ? AND user_id != ?", authorId, oldId).
		Where("user_id NOT IN (?)", subQuery).
		Order("RANDOM()").
		Limit(consts.One).
		Scan(&userId)

	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "", gorm.ErrRecordNotFound
	}

	return userId, nil
}
