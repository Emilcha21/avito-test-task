package pullrequest

import "avito-test/internal/models"

func (r *PullRequestRepoSQL) CreatePullRequest(model *models.PullRequest) error {
    result := r.db.Create(model)
    if result.Error != nil {
        return result.Error
    }

    return nil
}