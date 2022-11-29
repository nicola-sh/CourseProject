package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nicola-sh/CourseProject/Hike/sqlc"
)

type createRouteRequest struct {
	UserID      int32   `json:"user_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Location    string  `json:"location"`
	Destination float64 `json:"destination"`
	Height      float64 `json:"height"`
	Level       string  `json:"level", binding:"required", oneof=low middle high`
}

func (server *Server) createRoute(ctx *gin.Context) {
	var req createRouteRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateRouteParams{
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		Destination: req.Destination,
		Height:      req.Height,
		Level:       req.Level,
	}

	account, err := server.store.CreateRoute(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
