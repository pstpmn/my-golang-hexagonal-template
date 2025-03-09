package httpHandler

import (
	"github.com/labstack/echo/v4"
)

type (
	Response struct {
		Message    string `json:"message"`
		Result     any    `json:"result,omitempty"`
		StatusCode int    `json:"statusCode"`
	}
)

func ResponseError(statusCode int, message string, ctx echo.Context) error {
	return ctx.JSON(statusCode, &Response{
		Message:    message,
		StatusCode: statusCode,
	})
}

func ResponseSuccess(statusCode int, message string, result any, ctx echo.Context) error {
	return ctx.JSON(statusCode, &Response{
		Message:    message,
		Result:     result,
		StatusCode: statusCode,
	})
}
