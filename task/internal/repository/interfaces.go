package repository

import (
	"avito-test/internal/models"
	"time"
)

type ITeamRepo interface {
	CreateTeam(reqTeam *models.Team, reqUsers []models.User) error
	GetTeamByName(name string) (string, error)
}

type IUserRepo interface {
	GetUsersByTeam(name string) (*[]models.User, error)
	GetUserById(id string) (models.User, error)
	ChangeActiveStatusById(id string, isActive bool) error
	GetUsersToPr(teamName string, authorId string) ([]string, error)
	GetUsersByPr(prName string) ([]string, error)
	GetUserToPr(teamName string, authorId string, prId string, oldId string) (string, error)
	ChangeUserAssignee(prId string, userId string, newUserId string) error
}

type IPullRequestRepo interface {
	CreatePullRequest(req *models.PullRequest) error
	GetPullRequestById(id string) (*models.PullRequest, error)
	AddReviewers(usersId []string, prId string) error
	SetMergeStatus(prId string, mergedAt time.Time) error
	GetPullRequestByUser(userId string, pullRequestId string) (bool, error)
	GetPullRequestsByUser(userId string) ([]models.PullRequest, error)
}
