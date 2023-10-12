package gen

import "github.com/go-playground/validator/v10"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var Validate = validator.New()
var ValidateResponse = &Response{Code: 400, Message: "param validator error"}
