package cli

import (
	"transfeed/internal/app/store"
	"transfeed/internal/util"

	"github.com/spf13/cobra"
)

var dsUrl string

func init() {
	rootCmd.Flags().StringVarP(&dsUrl, "database-url", "d", util.ExcutePath()+"/transfeed.db", "database url")
	rootCmd.AddCommand(resetUserCmd)
}

var resetUserCmd = &cobra.Command{
	Use:   "reset",
	Short: "reset user password",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		store.InitDB(dsUrl)
		store.ResetUser(args[0], args[1])
	},
}
