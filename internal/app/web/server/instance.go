package server

import (
	"fmt"
	"transfeed/internal/app/web/api"
	"transfeed/internal/app/web/config"
	"transfeed/internal/app/web/view"
	"transfeed/internal/app/web/worker"
	"transfeed/internal/env"
	"transfeed/internal/util"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "transfeed/docs"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	Config = config.Config{
		JWTScrect: util.ShortUID(12),
	}
)

func MakeRouter(env *env.Env) *echo.Echo {

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods: []string{
			echo.GET,
			echo.HEAD,
			echo.PUT,
			echo.PATCH,
			echo.POST,
			echo.DELETE,
		},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	// 	TokenLookup: "header:X-XSRF-TOKEN",
	// }))

	// 全局中间件：为所有路由设置配置项
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", Config)
			return next(c)
		}
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	return e
}

func RunServer(env *env.Env) {
	e := MakeRouter(env)
	timer := worker.NewTimer(env.RefreshHours)
	publicGroup := e.Group(env.UrlPrefix)
	protectGroup := e.Group(env.UrlPrefix)
	protectGroup.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(Config.JWTScrect),
		Skipper: func(c echo.Context) bool {
			cPath := c.Path()
			switch cPath {
			case "/api/v1/user/login":
				return true
			default:
				return false
			}
		},
	}))

	if !env.DisableSwagger {
		publicGroup.GET("/swagger/*", echoSwagger.WrapHandler)
		fmt.Printf("swagger at:http://127.0.0.1:%d%s/swagger/index.html\n", env.Port, env.UrlPrefix)

	}
	timer.Start()

	api.UserAttach(protectGroup, publicGroup)
	api.FeedAttach(protectGroup, publicGroup)
	view.PageAttach(publicGroup)
	api.TranslatorAttach(protectGroup, publicGroup)
	fmt.Printf("app server at:http://127.0.0.1:%d%s\n", env.Port, env.UrlPrefix)
	e.Logger.Fatal(e.Start(fmt.Sprintf("127.0.0.1:%d", env.Port)))

}
