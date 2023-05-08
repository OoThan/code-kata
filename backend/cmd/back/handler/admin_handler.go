package handler

import (
	"loan-back-services/pkg/dto"
	"loan-back-services/pkg/logger"
	"loan-back-services/pkg/model"
	"loan-back-services/pkg/repository"
	"loan-back-services/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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
	// group.Use(middleware.AuthMiddleware(ctr.repo))

	group.POST("/list", ctr.listAdmin)
	group.POST("/add", ctr.addAdmin)
	group.POST("/edit", ctr.editAdmin)
	group.POST("/delete", ctr.deleteAdmin)
}

func (ctr *adminHandler) listAdmin(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.AdminListReq{}
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	list, total, err := ctr.repo.Admin.List(c.Request.Context(), req)
	if err != nil {
		res = utils.GenerateGormErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	data := gin.H{
		"list":  list,
		"total": total,
	}

	res = utils.GenerateSuccessResponse(data)
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *adminHandler) addAdmin(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.AdminAddReq{}
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		res = utils.GenerateInternalServerErrorResponse("Unable to hash password")
		c.JSON(200, res)
		c.Abort()
		return
	}
	req.Password = hashPassword

	admin := &model.Admin{}
	if err := copier.Copy(&admin, &req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}
	admin.IP = c.ClientIP()
	area, err := utils.GetArea(admin.IP)
	if err != nil {
		logger.Sugar.Error(err)
	}
	admin.Location = area

	err = ctr.repo.Admin.Create(c.Request.Context(), admin)
	if err != nil {
		res = utils.GenerateGormErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	res = utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *adminHandler) editAdmin(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.AdminEditReq{}
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	updateFields := &model.UpdateFields{
		Field: "id",
		Value: req.Id,
		Data:  map[string]any{},
	}
	if req.Password != "" {
		hashPassword, err := utils.HashPassword(req.Password)
		if err != nil {
			res = utils.GenerateInternalServerErrorResponse("Unable to hash password")
			c.JSON(200, res)
			c.Abort()
			return
		}
		updateFields.Data["password"] = hashPassword
	}
	updateFields.Data["username"] = req.Username
	updateFields.Data["email"] = req.Email
	updateFields.Data["status"] = req.Status

	if err := ctr.repo.Admin.Update(c.Request.Context(), updateFields); err != nil {
		res = utils.GenerateGormErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	res = utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *adminHandler) deleteAdmin(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.AdminDeleteReq{}
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	ids := utils.IdsIntToInCon(req.Ids)
	err := ctr.repo.Admin.Delete(c.Request.Context(), ids)
	if err != nil {
		res = utils.GenerateGormErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	res = utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}
