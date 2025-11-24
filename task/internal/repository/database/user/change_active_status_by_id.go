package user

import "avito-test/internal/models"

func (r *UserRepoSQL) ChangeActiveStatusById(id string, isActive bool) error {
	result := r.db.Model(&models.User{}).
		Where("user_id = ?", id).
		Update("is_active", isActive)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
