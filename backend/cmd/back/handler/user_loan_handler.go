package handler

import (
	"github.com/gin-gonic/gin"
	"loan-back-services/pkg/middleware"
	"loan-back-services/pkg/repository"
)

type userLoanHandler struct {
	R    *gin.Engine
	repo *repository.Repository
}

func newUserLoanHandler(h *Handler) *userLoanHandler {
	return &userLoanHandler{
		R:    h.R,
		repo: h.repo,
	}
}

func (ctr *userLoanHandler) Register() {
	group := ctr.R.Group("/api/user-loan")
	group.Use(middleware.AuthMiddleware(ctr.repo))
}
