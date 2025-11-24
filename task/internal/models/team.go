package models

type Team struct {
	TeamName string `gorm:"column:team_name;primaryKey"`
}
