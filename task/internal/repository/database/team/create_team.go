package team

import (
	"avito-test/consts"
	"avito-test/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *TeamRepoSQL) CreateTeam(reqTeam *models.Team, reqUsers []models.User) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(reqTeam).Error
		if err != nil {
			return err
		}

		if len(reqUsers) == 0 {
			return nil
		}

		for i := range reqUsers {
			reqUsers[i].TeamName = reqTeam.TeamName
		}

		err = tx.
			Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"username", "is_active", "team_name"}),
			}).
			CreateInBatches(&reqUsers, consts.OneHundred).Error
		if err != nil {
			return err
		}

		return nil
	})
}
