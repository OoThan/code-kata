package handler

import (
	"github.com/gin-gonic/gin"
	"loan-back-services/pkg/middleware"
	"loan-back-services/pkg/repository"
)

type adminHandler struct {
	R    *gin.Engine
	repo *repository.Repository
}

func newAdminHandler(h *Handler) *adminHandler {
	return &adminHandler{
		R:    h.R,
		repo: h.repo,
	}
}

func (ctr *adminHandler) Register() {
	group := ctr.R.Group("/api/admin")
	group.Use(middleware.AuthMiddleware(ctr.repo))
}
