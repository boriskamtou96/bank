package api

import (
	"fmt"
	db "simplebank/db/sqlc"
	"simplebank/token"
	"simplebank/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     utils.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("token.NewPasetoMaker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("currency", validCurrency); err != nil {
			panic(err)
		}
	}

	authRoutes := router.Group("/", authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts/", server.listAccounts)
	authRoutes.PUT("/accounts/", server.updateAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)

	authRoutes.POST("/transfers/", server.createTransfer)

	router.POST("/users/", server.createUser)
	router.POST("/users/login/", server.loginUser)

	server.router = router
	return server, nil
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
