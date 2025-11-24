package models

import "time"

type PullRequest struct {
	PullRequestId     string     `gorm:"column:pull_request_id;primaryKey"`
	PullRequestName   string     `gorm:"column:pull_request_name;not null"`
	Status            string     `gorm:"column:status;not null;default:'OPEN'"`
	AuthorId          string     `gorm:"column:author_id;not null"`
	NeedMoreReviewers bool       `gorm:"column:need_more_reviewers;not null;default:true"`
	CreatedAt         time.Time  `gorm:"column:created_at;default:now()"`
	MergedAt          *time.Time `gorm:"column:merged_at"`
}
