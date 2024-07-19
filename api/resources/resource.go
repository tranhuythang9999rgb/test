package resources

import (
	"github.com/gofiber/fiber/v2"
)

type response struct {
	HTTPCode int         `json:"http_code"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"` // omitempty will omit the field if it's empty
}

func ResponseSuccess(c *fiber.Ctx, data ...interface{}) error {
	var respData interface{}
	if len(data) > 0 {
		respData = data[0]
	} else {
		respData = nil
	}
	resp := response{
		HTTPCode: 200,
		Code:     0,
		Message:  "success",
		Data:     respData,
	}
	return c.JSON(resp)
}
