package utils

import (
	"loan-back-services/pkg/dto"
	"net/http"
)

func GenerateAuthErrorResponse(msg string) *dto.ResponseObject {
	res := &dto.ResponseObject{}
	res.ErrCode = 401
	res.ErrMsg = msg
	res.HttpStatusCode = http.StatusUnauthorized
	return res
}
