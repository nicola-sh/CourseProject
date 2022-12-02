package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

	db "github.com/nicola-sh/CourseProject/Hike/sqlc"
	"github.com/nicola-sh/CourseProject/Hike/token"
	"github.com/nicola-sh/CourseProject/Hike/util"
)

// Server serves HTTP requests for hike service
type Server struct {
	config     util.Config
	tokenMaker token.Maker
	store      *db.Store
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("Can't create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.listUser)

	server.router = router
	return server, nil
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
