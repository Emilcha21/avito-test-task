package user

import "avito-test/internal/models"

func (r *UserRepoSQL) ChangeUserAssignee(prId string, userId string, newUserId string) error {
	result := r.db.Model(&models.PRUsers{}).
		Where("pull_request_id = ? AND user_id = ?", prId, userId).
		Update("user_id", newUserId)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
