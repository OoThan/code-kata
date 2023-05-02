package dto

type ResponseObject struct {
	ErrCode uint64 `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
	Data    any    `json:"data"`
}
