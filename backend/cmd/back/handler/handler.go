package handler

import (
	"github.com/gin-gonic/gin"
	"loan-back-services/pkg/ds"
	"loan-back-services/pkg/middleware"
)

type Handler struct {
	R *gin.Engine
}

type HConfig struct {
	R  *gin.Engine
	DS *ds.DataSource
}

func NewHandler(c *HConfig) *Handler {
	return &Handler{
		R: c.R,
	}
}

func (h *Handler) Register() {
	// middleware
	h.R.Use(middleware.Cors())
}

func (h *Handler) Destroy() {
}
