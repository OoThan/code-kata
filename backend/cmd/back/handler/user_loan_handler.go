package handler

import (
	"loan-back-services/pkg/dto"
	"loan-back-services/pkg/middleware"
	"loan-back-services/pkg/repository"
	"loan-back-services/pkg/utils"

	"github.com/gin-gonic/gin"
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

func (ctr *userLoanHandler) usernameFilter(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.UsernameFilterListReq{}
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	list, err := ctr.repo.User.UsernameFilterList(c.Request.Context(), req)
	if err != nil {
		res = utils.GenerateGormErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	res = utils.GenerateSuccessResponse(gin.H{"list": list})
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *userLoanHandler) packageNoFilter(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.PackageNameFilterListReq{}
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	list, err := ctr.repo.LoanPackage.PackageNoFilterList(c.Request.Context(), req)
	if err != nil {
		res = utils.GenerateGormErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	res = utils.GenerateSuccessResponse(gin.H{"list": list})
	c.JSON(res.HttpStatusCode, res)
}
