package models

type PRUsers struct {
	ID            uint   `gorm:"column:id;primaryKey;autoIncrement"`
	PullRequestId string `gorm:"column:pull_request_id;not null"`
	UserId        string `gorm:"column:user_id;not null"`
}
