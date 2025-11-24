package user

import "avito-test/internal/models"

func (r *UserRepoSQL) GetUsersToPr(teamName string, authorId string) ([]string, error) {
	var userIds []string

	result := r.db.Model(&models.User{}).
		Select("user_id").
		Where("team_name = ? AND is_active = TRUE AND user_id != ?", teamName, authorId).
		Order("RANDOM()").
		Limit(2).
		Pluck("user_id", &userIds)

	if result.Error != nil {
		return nil, result.Error
	}

	return userIds, nil
}
