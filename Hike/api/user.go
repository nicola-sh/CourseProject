package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nicola-sh/CourseProject/Hike/sqlc"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Age      int32  `json:"age" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required, oneof=admin photographer hiker"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Age:      req.Age,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	}

	account, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
