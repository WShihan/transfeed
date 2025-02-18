package main

import "transfeed/internal/cli"

var (
	Version   string
	BuildTime string
	Commit    string
	Author    = "Wangshihan"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8090
// @securityDefinitions.apikey  JWT
// @in                          header
// @name                        Authorization
// @description                 Bearer Token 格式：`Bearer {token}`
// @Param   Authorization  header  string  false  "Bearer Token"
// @name Authorization
// @description JWT Authorization header
func main() {
	cli.Parse(cli.Injection{
		Version:   Version,
		BuildTime: BuildTime,
		Commit:    Commit,
		Author:    Author,
	})
}
