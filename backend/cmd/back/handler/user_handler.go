package handler

import (
	"loan-back-services/pkg/dto"
	"loan-back-services/pkg/model"
	"loan-back-services/pkg/repository"
	"loan-back-services/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type userHandler struct {
	R    *gin.Engine
	repo *repository.Repository
}

func newUserHandler(h *Handler) *userHandler {
	return &userHandler{
		R:    h.R,
		repo: h.repo,
	}
}

func (ctr *userHandler) Register() {
	group := ctr.R.Group("/api/user")
	// group.Use(middleware.AuthMiddleware(ctr.repo))

	group.POST("/list", ctr.listUser)
	group.POST("/add", ctr.addUser)
	group.POST("/edit", ctr.editUser)
	group.POST("/delete", ctr.deleteUser)
}

func (ctr *userHandler) listUser(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.UserListReq{}
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	list, total, err := ctr.repo.User.List(c.Request.Context(), req)
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

func (ctr *userHandler) addUser(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.UserAddReq{}
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	user := &model.User{}
	if err := copier.Copy(&user, &req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	if err := ctr.repo.User.Create(c.Request.Context(), user); err != nil {
		res = utils.GenerateGormErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	res = utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *userHandler) editUser(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.UserEditReq{}
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
	updateFields.Data["username"] = req.Username
	updateFields.Data["user_nrc"] = req.UserNRC
	updateFields.Data["user_phone_number"] = req.UserPhoneNumber
	updateFields.Data["reference_user_name"] = req.ReferenceUserName
	updateFields.Data["street"] = req.Street
	updateFields.Data["city"] = req.City
	updateFields.Data["region"] = req.Region

	if err := ctr.repo.User.Update(c.Request.Context(), updateFields); err != nil {
		res = utils.GenerateGormErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	res = utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)

}

func (ctr *userHandler) deleteUser(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.UserDeleteReq{}
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	ids := utils.IdsIntToInCon(req.Ids)
	err := ctr.repo.User.Delete(c.Request.Context(), ids)
	if err != nil {
		res = utils.GenerateGormErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	res = utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}
