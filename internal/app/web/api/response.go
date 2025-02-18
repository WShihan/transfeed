package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Status  bool        `json:"status"`                    // 状态
	Code    int         `json:"code" example:"200"`        // 状态码
	Message string      `json:"message" example:"Success"` // 消息
	Data    interface{} `json:"data"`                      // 数据
}

func ApiSuccess(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Code:    200,
		Status:  true,
		Message: "success",
		Data:    data,
	})
}

func ApiFailed(c echo.Context, code int, msg string) error {
	return c.JSON(code, Response{
		Code:    code,
		Status:  false,
		Message: msg,
		Data:    nil,
	})
}
