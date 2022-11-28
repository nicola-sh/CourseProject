package api

import "github.com/gin-gonic/gin"

// Server serves HTTP requests for hike service
type Server struct {
	router *gin.Engine
}
