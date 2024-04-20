package api

import (
	"dhack-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getStatRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type getStatListRequest struct {
	Offset int32 `form:"offset"`
	Limit  int32 `form:"limit" binding:"required,min=1,max=20"`
	UserID int32 `form:"user_id" binding:"required"`
}

type createStatRequest struct {
	CaloriesBurned int32   `json:"calories_burned" binding:"required"`
	Rpm            float32 `json:"rpm" binding:"required"`
	Duration       int32   `json:"duration" binding:"required"`
	Score          float32 `json:"score" binding:"required"`
	ChallengeID    int32   `json:"challenge_id" binding:"required"`
	UserID         int32   `json:"user_id" binding:"required"`
}

type updateStatRequest struct {
	CaloriesBurned int32   `json:"calories_burned" binding:"required"`
	Rpm            float32 `json:"rpm" binding:"required"`
	Duration       int32   `json:"duration" binding:"required"`
	Score          float32 `json:"score" binding:"required"`
	ChallengeID    int32   `json:"challenge_id" binding:"required"`
	UserID         int32   `json:"user_id" binding:"required"`
}

func (server *Server) GetStatByID(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getStatRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	// Execute query.
	result, err := server.store.Queries.GetStat(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) GetStatListByUser(ctx *gin.Context) {

	// Check if request has parameters offset and limit for pagination.
	var req getStatListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	arg := model.ListStatsByUserParams{
		Offset: req.Offset,
		Limit:  req.Limit,
		UserID: req.UserID,
	}

	// Execute query.
	result, err := server.store.Queries.ListStatsByUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) CreateStat(ctx *gin.Context) {

	// Check if request has all required fields in JSON body.
	var req createStatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	arg := model.CreateStatParams{
		CaloriesBurned: req.CaloriesBurned,
		Rpm:            req.Rpm,
		Duration:       req.Duration,
		Score:          req.Score,
		ChallengeID:    req.ChallengeID,
		UserID:         req.UserID,
	}

	// Execute query.
	result, err := server.store.Queries.CreateStat(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (server *Server) UpdateStat(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var reqID getStatRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	// Check if request has all required fields in JSON body.
	var req updateStatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	arg := model.UpdateStatParams{
		CaloriesBurned: req.CaloriesBurned,
		Rpm:            req.Rpm,
		Duration:       req.Duration,
		Score:          req.Score,
		ChallengeID:    req.ChallengeID,
		UserID:         req.UserID,
	}

	// Execute query.
	result, err := server.store.Queries.UpdateStat(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (server *Server) DeleteStat(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getStatRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	// Execute query.
	err := server.store.Queries.DeleteStat(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (server *Server) GetWeeklyProgress(ctx *gin.Context) {

	var req getStatRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	// Execute query.
	result, err := server.store.Queries.GetWeekyProgress(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}
