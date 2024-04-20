package api

import (
	"dhack-api/model"
	"dhack-api/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

type getUserListRequest struct {
	Offset int32 `form:"offset"`
	Limit  int32 `form:"limit" binding:"required,min=1,max=20"`
}

type createUserRequest struct {
	Name      string          `json:"name" binding:"required"`
	Email     string          `json:"email" binding:"required"`
	Weight    int16           `json:"weight" binding:"required"`
	BirthDate util.CustomTime `json:"birth_date" binding:"required"`
}

type updateUserRequest struct {
	Name      string          `json:"name" binding:"required"`
	Email     string          `json:"email" binding:"required"`
	Weight    int16           `json:"weight" binding:"required"`
	BirthDate util.CustomTime `json:"birth_date" binding:"required"`
}

func (server *Server) GetUserByID(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Execute query.
	result, err := server.store.Queries.GetUser(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) GetUserList(ctx *gin.Context) {

	// Check if request has parameters offset and limit for pagination.
	var req getUserListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Execute query.
	result, err := server.store.Queries.ListUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (server *Server) CreateUser(ctx *gin.Context) {

	// Check if request has all required fields in JSON body.
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := model.CreateUserParams{
		Name:      req.Name,
		Email:     req.Email,
		Weight:    req.Weight,
		BirthDate: pgtype.Date{Time: time.Time(req.BirthDate), Valid: true},
	}

	// Execute query.
	result, err := server.store.Queries.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (server *Server) UpdateUser(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var reqID getUserRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Check if request has all required fields in JSON body.
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	arg := model.UpdateUserParams{
		Name:      req.Name,
		Email:     req.Email,
		Weight:    req.Weight,
		BirthDate: pgtype.Date{Time: time.Time(req.BirthDate), Valid: true},
		ID:        reqID.ID,
	}

	// Execute query.
	result, err := server.store.Queries.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (server *Server) DeleteUser(ctx *gin.Context) {

	// Check if request has ID field in URI.
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		ctx.Abort()
		return
	}

	// Execute query.
	err := server.store.Queries.DeleteUser(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
