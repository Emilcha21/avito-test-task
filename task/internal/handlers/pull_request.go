package handlers

import (
	"avito-test/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

//		@Summary      Create Pull Request
//		@Description  Create a new pull request. Requires admin authorization
//		@Tags         Pull Requests
//		@Accept       json
//		@Produce      json
//		@Param        request    body     dto.CreatePullRequestReq  true  "Pull request creation data"
//		@Success      201  {object}   dto.CreatePullRequestResp
//		@Failure      400  {object}  dto.ErrorResponse
//		@Failure      401  {object}  dto.ErrorResponse
//		@Failure      404  {object}  dto.ErrorResponse
//		@Failure      500  {object}  dto.ErrorResponse
//		@Router       /pullRequest/create [post]
//	 	@Security 	  BearerAuth
func (h *Handler) CreatePullRequest(c *gin.Context) {
	var req dto.CreatePullRequestReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	reviewersId, err := h.srv.CreatePullRequest(req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"pr": dto.CreatePullRequestResp{
		PullRequestId:   req.PullRequestId,
		PullRequestName: req.PullRequestName,
		AuthorId:        req.AuthorId,
		Status:          "OPEN",
		Reviewers:       reviewersId,
	}})
}

//		@Summary      Merge Pull Request
//		@Description  Update pull request status to MERGED. Requires admin authorization
//		@Tags         Pull Requests
//		@Accept       json
//		@Produce      json
//		@Param        request    body     dto.SetMergeStatusReq  true  "Merge status update data"
//		@Success      200  {object}   dto.PullRequestResp
//		@Failure      400  {object}  dto.ErrorResponse
//		@Failure      401  {object}  dto.ErrorResponse
//		@Failure      404  {object}  dto.ErrorResponse
//		@Failure      500  {object}  dto.ErrorResponse
//		@Router       /pullRequest/merge [patch]
//	 	@Security 	  BearerAuth
func (h *Handler) SetMergeStatus(c *gin.Context) {
	var req dto.SetMergeStatusReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	pullRequest, err := h.srv.SetMergeStatus(req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"pr": dto.PullRequestResp{
		PullRequestId:   pullRequest.PullRequestId,
		PullRequestName: pullRequest.PullRequestName,
		AuthorId:        pullRequest.AuthorId,
		Status:          pullRequest.Status,
		Reviewers:       pullRequest.Reviewers,
		MergedAt:        pullRequest.MergedAt,
	}})
}

//	 	@Summary      Reassign Reviewer
//		@Description  Reassign pull request reviewer to another user. Requires admin authorization
//		@Tags         Pull Requests
//		@Accept       json
//		@Produce      json
//		@Param        request    body     dto.ReassignReviewerReq  true  "Reviewer reassignment data"
//		@Success      200  {object}   dto.ReassignReviewerResp
//		@Failure      400  {object}  dto.ErrorResponse
//		@Failure      401  {object}  dto.ErrorResponse
//		@Failure      404  {object}  dto.ErrorResponse
//		@Failure      409  {object}  dto.ErrorResponse
//		@Failure      500  {object}  dto.ErrorResponse
//		@Router       /pullRequest/reassign [post]
//	 	@Security 	  BearerAuth
func (h *Handler) ReassignReviewer(c *gin.Context) {
	var req dto.ReassignReviewerReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	resp, err := h.srv.ReassignReviewer(req)
	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pr":          resp,
		"replaced_by": req.UserId})
}
