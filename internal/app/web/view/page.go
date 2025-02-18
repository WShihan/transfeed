package view

import (
	"embed"
	"fmt"
	"net/http"
	"strings"
	"text/template"
	"transfeed/internal/app/web/assets"
	"transfeed/internal/util"

	"github.com/labstack/echo/v4"
)

func AssetsFinder(c echo.Context) error {
	assettype := c.Param("assettype")
	fileName := c.Param("filename")
	assetDir := fmt.Sprintf("static/%s/%s", assettype, fileName)
	var fs embed.FS
	if strings.HasSuffix(fileName, ".js") || strings.HasSuffix(fileName, ".map") {
		fs = assets.JS
		c.Response().Header().Set("Content-Type", "application/javascript")
	} else if strings.HasSuffix(fileName, ".png") {
		fs = assets.IMG
		c.Response().Header().Set("Content-Type", "image/x-icon")

	} else if strings.HasSuffix(fileName, ".ico") {
		fs = assets.IMG
		c.Response().Header().Set("Content-Type", "image/x-icon")

	} else {
		fs = assets.CSS
		c.Response().Header().Set("Content-Type", "text/css")

	}
	content, err := fs.ReadFile(assetDir)
	if err != nil {
		util.Logger.Error(err.Error())
		return c.String(http.StatusNotFound, err.Error())
	}
	c.Response().Write(content)
	return nil
}

func PageIndex(c echo.Context) error {
	t, _ := template.ParseFS(assets.HTML, "index.html")
	t.Execute(c.Response(), nil)
	return nil
}

func PageAttach(gp *echo.Group) {
	gp.GET("/", PageIndex)
	gp.GET("/static/:assettype/:filename", AssetsFinder)
}
