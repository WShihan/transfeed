package util

import "github.com/labstack/echo/v4"

type CustomEcho struct {
	BaseURL string // 根路由
	Echo    *echo.Echo
}

// 带前缀的 GET 方法
func (r *CustomEcho) GET(path string, handle echo.HandlerFunc) {
	r.Echo.GET(r.BaseURL+path, handle)
}

// 带前缀的 POST 方法
func (r *CustomEcho) POST(path string, handle echo.HandlerFunc) {
	r.Echo.POST(r.BaseURL+path, handle)
}

// 带前缀的 PUT 方法
func (r *CustomEcho) PUT(path string, handle echo.HandlerFunc) {
	r.Echo.PUT(r.BaseURL+path, handle)
}

// 带前缀的 PATCH 方法
func (r *CustomEcho) PATCH(path string, handle echo.HandlerFunc) {
	r.Echo.PATCH(r.BaseURL+path, handle)
}

// 带前缀的 DELETE 方法
func (r *CustomEcho) DELETE(path string, handle echo.HandlerFunc) {
	r.Echo.DELETE(r.BaseURL+path, handle)
}

// 带前缀的 OPTIONS 方法
func (r *CustomEcho) OPTIONS(path string, handle echo.HandlerFunc) {
	r.Echo.OPTIONS(r.BaseURL+path, handle)
}
