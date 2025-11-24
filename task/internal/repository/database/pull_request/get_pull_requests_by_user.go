package pullrequest

import "avito-test/internal/models"

func (r *PullRequestRepoSQL) GetPullRequestsByUser(userId string) ([]models.PullRequest, error) {
	var pullRequests []models.PullRequest

	result := r.db.
		Table("pull_request").
		Joins("JOIN pr_users ON pr_users.pull_request_id = pull_request.pull_request_id").
		Where("pr_users.user_id = ?", userId).
		Find(&pullRequests)

	if result.Error != nil {
		return nil, result.Error
	}

	return pullRequests, nil
}
