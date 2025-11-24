package team

import (
	"avito-test/internal/models"
	"errors"

	"gorm.io/gorm"
)

func (r *TeamRepoSQL) GetTeamByName(name string) (string, error) {
	var team models.Team

	result := r.db.
		Select("team_name").
		Where("team_name = ?", name).
		First(&team)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", gorm.ErrRecordNotFound
		}
		return "", result.Error
	}

	return team.TeamName, nil
}
