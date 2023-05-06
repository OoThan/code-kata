package utils

import (
	"fmt"
	"loan-back-services/pkg/dto"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func GenerateAuthErrorResponse(msg string) *dto.ResponseObject {
	res := &dto.ResponseObject{}
	res.ErrCode = 401
	res.ErrMsg = msg
	res.HttpStatusCode = http.StatusUnauthorized
	return res
}

func GenerateBindingErrorResponse(err error) *dto.ResponseObject {
	res := &dto.ResponseObject{}
	res.ErrCode = http.StatusUnprocessableEntity
	res.ErrMsg = GenerateValidationErrorMessage(err)
	res.HttpStatusCode = http.StatusUnprocessableEntity
	return res
}

func GenerateInternalServerErrorResponse(msg string) *dto.ResponseObject {
	res := &dto.ResponseObject{}
	res.ErrCode = http.StatusInternalServerError
	res.ErrMsg = msg
	res.HttpStatusCode = http.StatusInternalServerError
	return res
}

func GenerateDisableUserResponse(msg string) *dto.ResponseObject {
	res := &dto.ResponseObject{}
	res.ErrCode = http.StatusBadRequest
	res.ErrMsg = msg
	res.HttpStatusCode = http.StatusBadRequest
	return res
}

func GenerateGormErrorResponse(err error) *dto.ResponseObject {
	res := &dto.ResponseObject{}
	res.ErrMsg = err.Error()
	if IsErrNotFound(err) {
		res.ErrCode = 400
		res.HttpStatusCode = http.StatusBadRequest
		return res
	}

	if IsDuplicate(err) {
		fields := strings.Split(err.Error(), ".")
		field := fields[len(fields)-1]
		res.ErrCode = 400
		msg := "is already existed"
		res.ErrMsg = fmt.Sprintf("%v %s", strings.Trim(field, "'"), msg)
		res.HttpStatusCode = http.StatusBadRequest
		return res
	}
	res.ErrCode = 500
	res.HttpStatusCode = http.StatusInternalServerError
	return res
}

func GenerateSuccessResponse(data any) *dto.ResponseObject {
	res := &dto.ResponseObject{}
	res.ErrCode = 0
	res.ErrMsg = "SUCCESS"
	res.Data = data
	res.HttpStatusCode = http.StatusOK
	return res
}

func msgForTag(fa validator.FieldError) string {
	switch fa.Tag() {
	case "required":
		return fmt.Sprintf("%v field is required.", fa.Field())
	case "oneof":
		return fmt.Sprintf("%v field must be one of %v", fa.Field(), fa.Param())
	case "email":
		return "Invalid email."
	}

	return fa.Field()
}

func GenerateValidationErrorMessage(err error) string {
	if vErr, ok := err.(validator.ValidationErrors); ok {
		ErrMsg := ""
		for _, fieldErr := range vErr {
			ErrMsg += msgForTag(fieldErr)
		}
		return ErrMsg
	}

	return err.Error()
}
