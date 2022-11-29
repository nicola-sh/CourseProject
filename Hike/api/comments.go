package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nicola-sh/CourseProject/Hike/sqlc"
)

type createCommentRequest struct {
	CommentText string `json:"comment_text"`
}

func (server *Server) createComment(ctx *gin.Context) {
	var req createCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCommentParams{
		CommentText: req.CommentText,
	}

	account, err := server.store.CreateComment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
