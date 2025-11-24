package pullrequest

import "avito-test/internal/models"

func (r *PullRequestRepoSQL) AddReviewers(userIds []string, prId string) error {
	if len(userIds) == 0 {
		return nil
	}

	var prUsers []models.PRUsers
	for _, userId := range userIds {
		prUsers = append(prUsers, models.PRUsers{
			PullRequestId: prId,
			UserId:        userId,
		})
	}

	result := r.db.Create(&prUsers)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
