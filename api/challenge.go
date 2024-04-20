package api

import (
	"dhack-api/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type getChallengeRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type getChallengeListRequest struct {
	Offset int32 `form:"offset"`
	Limit  int32 `form:"limit" binding:"required,min=1,max=20"`
}

type createChallengeRequest struct {
	Name     string          `json:"name" binding:"required"`
	Steps    json.RawMessage `json:"steps" binding:"required"`
	FileName string          `json:"file_name" binding:"required"`
	Duration int32           `json:"duration" binding:"required"`
}

type updateChallengeRequest struct {
	Name     string          `json:"name" binding:"required"`
	Steps    json.RawMessage `json:"steps" binding:"required"`
	FileName string          `json:"file_name" binding:"required"`
	Duration int32           `json:"duration" binding:"required"`
}

type getChallengeResponse struct {
	ID       int32           `json:"id"`
	Name     string          `json:"name"`
	Steps    json.RawMessage `json:"steps"`
	FileName string          `json:"file_name"`
	Duration int32           `json:"duration"`
}

func (server *Server) GetChallengeByID(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getChallengeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	// Execute query.
	result, err := server.store.Queries.GetChallenge(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) GetChallengeList(ctx *gin.Context) {

	// Check if request has parameters offset and limit for pagination.
	var req getChallengeListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	arg := model.ListChallengesParams{
		Offset: req.Offset,
		Limit:  req.Limit,
	}

	// Execute query.
	result, err := server.store.Queries.ListChallenges(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	var response []getChallengeResponse
	for _, elem := range result {
		response = append(response, getChallengeResponse{
			ID:       elem.ID,
			Name:     elem.Name,
			Steps:    elem.Steps,
			FileName: elem.FileName,
			Duration: elem.Duration,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

func (server *Server) CreateChallenge(ctx *gin.Context) {

	// Check if request has all required fields in JSON body.
	var req createChallengeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println("We have an error: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	arg := model.CreateChallengeParams{
		Name:     req.Name,
		Steps:    req.Steps,
		FileName: req.FileName,
		Duration: req.Duration,
	}

	// Execute query.
	result, err := server.store.Queries.CreateChallenge(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	response := getChallengeResponse{
		ID:       result.ID,
		Name:     result.Name,
		Steps:    result.Steps,
		FileName: result.FileName,
		Duration: result.Duration,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (server *Server) UpdateChallenge(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var reqID getChallengeRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	// Check if request has all required fields in JSON body.
	var req updateChallengeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	arg := model.UpdateChallengeParams{
		Name:     req.Name,
		Steps:    req.Steps,
		FileName: req.FileName,
		Duration: req.Duration,
	}

	// Execute query.
	result, err := server.store.Queries.UpdateChallenge(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	response := getChallengeResponse{
		ID:       result.ID,
		Name:     result.Name,
		Steps:    result.Steps,
		FileName: result.FileName,
		Duration: result.Duration,
	}

	ctx.JSON(http.StatusCreated, response)
}

func (server *Server) DeleteChallenge(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getChallengeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	// Execute query.
	err := server.store.Queries.DeleteChallenge(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
