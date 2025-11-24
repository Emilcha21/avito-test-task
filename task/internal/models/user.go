package models

type User struct {
	UserId   string `gorm:"column:user_id;primaryKey"`
	Username string `gorm:"column:username;not null"`
	IsActive bool   `gorm:"column:is_active;not null;default:true"`
	TeamName string `gorm:"column:team_name;not null"`
}
