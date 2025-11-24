package errors

import "errors"

var (
	ErrTeamExists     = errors.New("team already exist")
	ErrTeamNotFound   = errors.New("team not found")
	ErrUserNotFound   = errors.New("user not found")
	ErrPrNotFound     = errors.New("pull request not found")
	ErrInternalServer = errors.New("internal server error")
	ErrPrExists       = errors.New("pull request already exist")
	ErrPrMerged       = errors.New("cannot reassign on merged PR")
	ErrNotAssigned    = errors.New("reviewer is not assigned to this PR")
	ErrNoCandidate    = errors.New("no active replacement candidate in team")
	ErrNoTeamQuery    = errors.New("query parameter need to be team_name")
	ErrNoUserQuery    = errors.New("query parameter need to be user_id")
	ErrMissingHeader  = errors.New("missing authorization header")
	ErrInvalidFormat  = errors.New("invalid authorization header format")
	ErrUnauthorized   = errors.New("invalid token")
)
