package handler

import (
	"github.com/gin-gonic/gin"
	"loan-back-services/pkg/ds"
	"loan-back-services/pkg/middleware"
	"loan-back-services/pkg/repository"
)

type Handler struct {
	R    *gin.Engine
	repo *repository.Repository
}

type HConfig struct {
	R  *gin.Engine
	DS *ds.DataSource
}

func NewHandler(c *HConfig) *Handler {
	repo := repository.NewRepository(&repository.RepoConfig{
		DS: c.DS,
	})
	return &Handler{
		R:    c.R,
		repo: repo,
	}
}

func (h *Handler) Register() {
	// middleware
	h.R.Use(middleware.Cors())

	// auth handler
	authHandler := newAuthHandler(h)
	authHandler.Register()

	// admin handler
	adminHandler := newAdminHandler(h)
	adminHandler.Register()

	// user handler
	userHandler := newUserHandler(h)
	userHandler.Register()

	// loan package handler
	loanPackageHandler := newLoanPackageHandler(h)
	loanPackageHandler.Register()
}

func (h *Handler) Destroy() {
}
