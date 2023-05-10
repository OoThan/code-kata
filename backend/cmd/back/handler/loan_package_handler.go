package handler

import (
	"loan-back-services/pkg/dto"
	"loan-back-services/pkg/middleware"
	"loan-back-services/pkg/model"
	"loan-back-services/pkg/repository"
	"loan-back-services/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type loanPackageHandler struct {
	R    *gin.Engine
	repo *repository.Repository
}

func newLoanPackageHandler(h *Handler) *loanPackageHandler {
	return &loanPackageHandler{
		R:    h.R,
		repo: h.repo,
	}
}

func (ctr *loanPackageHandler) Register() {
	group := ctr.R.Group("/api/loan-pkg")
	group.Use(middleware.AuthMiddleware(ctr.repo))

	group.POST("/list", ctr.listLoanPkg)
	group.POST("/list-log", ctr.listLoanPkgLog)
	group.POST("/add", ctr.addLoanPkg)
	group.POST("/edit", ctr.editLoanPkg)
	group.POST("/delete", ctr.deleteLoanPkg)
}

func (ctr *loanPackageHandler) listLoanPkg(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.LoanPackageListReq{}
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	list, total, err := ctr.repo.LoanPackage.List(c.Request.Context(), req)
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

func (ctr *loanPackageHandler) listLoanPkgLog(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.LoanPackageLogListReq{}
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	list, total, err := ctr.repo.LoanPackage.LogList(c.Request.Context(), req)
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

func (ctr *loanPackageHandler) addLoanPkg(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.LoanPackageAddReq{}
	admin := c.MustGet("admin").(*model.Admin)
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	loanPkg := &model.LoanPackage{}
	loanPkg.Creator = admin.Id
	if err := copier.Copy(&loanPkg, &req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	err := ctr.repo.LoanPackage.Create(c.Request.Context(), loanPkg)
	if err != nil {
		res = utils.GenerateGormErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	res = utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *loanPackageHandler) editLoanPkg(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.LoanPackageEditReq{}
	admin := c.MustGet("admin").(*model.Admin)
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
	updateFields.Data["package_no"] = req.PackageNo
	updateFields.Data["creator"] = admin.Id
	updateFields.Data["amount"] = req.Amount
	updateFields.Data["percent"] = req.Percent

	err := ctr.repo.LoanPackage.Update(c.Request.Context(), updateFields)
	if err != nil {
		res = utils.GenerateGormErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	res = utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *loanPackageHandler) deleteLoanPkg(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.LoanPackageDeleteReq{}
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	ids := utils.IdsIntToInCon(req.Ids)
	err := ctr.repo.LoanPackage.Delete(c.Request.Context(), ids)
	if err != nil {
		res = utils.GenerateGormErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	res = utils.GenerateSuccessResponse(nil)
	c.JSON(res.HttpStatusCode, res)
}
