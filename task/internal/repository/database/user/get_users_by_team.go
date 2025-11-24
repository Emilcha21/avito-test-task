package user

import "avito-test/internal/models"

func (r *UserRepoSQL) GetUsersByTeam(name string) (*[]models.User, error) {
	var users []models.User

	result := r.db.
		Select("user_id, username, is_active").
		Where("team_name = ?", name).
		Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return &users, nil
}
