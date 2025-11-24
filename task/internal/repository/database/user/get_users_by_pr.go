package user

import "avito-test/internal/models"

func (r *UserRepoSQL) GetUsersByPr(prId string) ([]string, error) {
	var userIds []string

	result := r.db.Model(&models.PRUsers{}).
		Where("pull_request_id = ?", prId).
		Pluck("user_id", &userIds)

	if result.Error != nil {
		return nil, result.Error
	}

	return userIds, nil
}
