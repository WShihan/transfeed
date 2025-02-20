package cli

import (
	"transfeed/internal/app/store"
	"transfeed/internal/app/web/server"
	"transfeed/internal/env"
	"transfeed/internal/util"

	"github.com/spf13/cobra"
)

var (
	port                 int
	databaseURL          string
	defaultAdminName     string
	defaultAdminPassword string
	urlPrefix            string
	disableSwagger       bool
	refreshHours         int
)

func init() {
	serveCmd.Flags().IntVarP(&port, "port", "p", 8090, "port to listen on")
	serveCmd.Flags().StringVarP(&databaseURL, "database-url", "d", util.ExcutePath()+"/transfeed.db", "database url")
	serveCmd.Flags().StringVarP(&defaultAdminName, "admin-name", "a", "admin", "default admin name")
	serveCmd.Flags().StringVarP(&defaultAdminPassword, "admin-password", "w", "admin1234", "default admin password")
	serveCmd.Flags().StringVarP(&urlPrefix, "url-prefix", "u", "", "url prefix")
	serveCmd.Flags().BoolVarP(&disableSwagger, "disable-swagger", "s", false, "disable swagger")
	serveCmd.Flags().IntVarP(&refreshHours, "refresh-hours", "h", 4, "refresh hours")

	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start a webserver instance",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := env.Env{
			Port:             port,
			DataBaseURL:      databaseURL,
			DefaultAdminName: defaultAdminName,
			DefaultAdminPass: defaultAdminPassword,
			UrlPrefix:        urlPrefix,
			DisableSwagger:   disableSwagger,
			Version:          Inject.Version,
			RefreshHours:     refreshHours,
		}
		util.InitLogger()
		store.InitDB(cfg.DataBaseURL)
		store.InitAdmin(cfg.DefaultAdminName, cfg.DefaultAdminPass)
		server.RunServer(&cfg)
	},
}
