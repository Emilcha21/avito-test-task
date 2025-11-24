package user

import (
	"avito-test/internal/models"
	"errors"

	"gorm.io/gorm"
)

func (r *UserRepoSQL) GetUserById(id string) (models.User, error) {
	var model models.User

	result := r.db.
		Select("user_id, username, is_active, team_name").
		Where("user_id = ?", id).
		First(&model)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return models.User{}, result.Error
		}
		return models.User{}, result.Error
	}

	return model, nil
}
