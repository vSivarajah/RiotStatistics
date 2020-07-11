package utils

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}
