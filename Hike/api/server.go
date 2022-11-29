package api

import (
	"go/token"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for hike service
type Server struct {
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer()
