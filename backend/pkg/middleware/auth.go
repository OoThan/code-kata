package middleware

import (
	"loan-back-services/conf"
	"loan-back-services/pkg/dto"
	"loan-back-services/pkg/repository"
	"loan-back-services/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AccessToken string `header:"Authorization"`
}

func AuthMiddleware(r *repository.Repository) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := &authHeader{}
		res := &dto.ResponseObject{}
		if err := c.ShouldBindHeader(&h); err != nil {
			res = utils.GenerateAuthErrorResponse("Must provide `Authorization` header in format of `Bearer {token}`")
			c.JSON(200, res)
			c.Abort()
			return
		}

		accessToken := strings.Split(h.AccessToken, "Bearer ")
		if len(accessToken) != 2 {
			res = utils.GenerateAuthErrorResponse("Permission Denied!")
			c.JSON(200, res)
			c.Abort()
			return
		}

		// validate access token
		accessTokenClaim, err := utils.ValidateAccessToken(accessToken[1], conf.Rsa().PublicKey)
		if err != nil {
			res = utils.GenerateAuthErrorResponse("Permission Denied!")
			c.JSON(200, res)
			c.Abort()
			return
		}

		ctx := c.Request.Context()
		if accessTokenClaim.Admin == nil {
			res = utils.GenerateAuthErrorResponse("Permission Denied!")
			c.JSON(200, res)
			c.Abort()
			return
		}

		admin, err := r.Admin.FindByField(ctx, "id", accessTokenClaim.Admin.Id)
		if err != nil {
			res = utils.GenerateAuthErrorResponse("Permission Denied!")
			c.JSON(200, res)
			c.Abort()
			return
		}

		c.Set("admin", admin)
		//c.Set("DB", r.DS.DB)
		c.Next()
	}
}
