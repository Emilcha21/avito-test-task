package pullrequest

func (r *PullRequestRepoSQL) GetPullRequestByUser(userId string, pullRequestId string) (bool, error) {
	var count int64

	result := r.db.Table("pr_users").
		Where("user_id = ? AND pull_request_id = ?", userId, pullRequestId).
		Count(&count)

	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}
