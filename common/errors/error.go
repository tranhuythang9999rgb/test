package errors

import (
	"fmt"
)

type Error interface {
	error
	GetHttpCode() int
	GetCode() int
	GetMessage() string
}

type CustomError struct {
	HttpCode int    `json:"http_code"`
	Code     int    `json:"code"`
	Message  string `json:"message"`
}

type CustomErrorMessage struct {
	Message string `json:"message"`
}

func (c CustomErrorMessage) GetMessage() string {
	return c.Message
}
func (c CustomError) GetHttpCode() int {
	return c.HttpCode
}

func (c CustomError) GetCode() int {
	return c.Code
}

func (c CustomError) GetMessage() string {
	return c.Message
}

func (c CustomError) Error() string {
	return fmt.Sprintf("[%v] %s", c.Code, c.Message)
}

func NewCustomHttpError(httpCode int, code int, message string) *CustomError {
	return &CustomError{
		HttpCode: httpCode,
		Code:     code,
		Message:  message,
	}
}
func NewCustomError(message string) *CustomErrorMessage {
	return &CustomErrorMessage{
		Message: message,
	}
}
