package cli

import (
	"flag"
	"fmt"
	"transfeed/internal/app/store"
	"transfeed/internal/app/web/server"
	"transfeed/internal/env"
	"transfeed/internal/util"
)

func Parse(injection Injection) {
	port := flag.Int("port", 8090, "启动端口，默认8090")
	adminName := flag.String("admin-name", "admin", "初始化管理员用户名，默认”admin“")
	adminPassword := flag.String("admin-password", "admin1234", "初始化管理员密码，默认“admin1234”")
	dataBaseURL := flag.String("database-url", util.ExcutePath()+"/transfeed.db", "数据库路径，默认当前目录")
	urlPrefix := flag.String("url-prefix", "", "路由前缀，默认无")
	disableSwagger := flag.Bool("disable-swagger", true, "是否关闭swagger，默认关闭")
	showVersion := flag.Bool("version", false, "显示版本信息")
	RefreshHours := flag.Int("refresh-hours", 4, "刷新间隔小时 默认4小时")
	flag.Parse()

	cfg := env.Env{
		Port:             *port,
		DataBaseURL:      *dataBaseURL,
		DefaultAdminName: *adminName,
		DefaultAdminPass: *adminPassword,
		UrlPrefix:        *urlPrefix,
		DisableSwagger:   *disableSwagger,
		Version:          injection.Version,
		RefreshHours:     *RefreshHours,
	}

	if *showVersion {
		fmt.Printf("author:%s\nversion:%s\ncommit:%s\nbuild time:%s\n", injection.Author, injection.Version, injection.Commit, injection.BuildTime)
		return
	}
	util.InitLogger()
	store.InitDB(*dataBaseURL)
	store.InitAdmin(*adminName, *adminPassword)
	store.ResetUser(*adminName, *adminPassword)
	server.RunServer(&cfg)
}
