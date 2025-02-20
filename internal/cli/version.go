package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version infomarion",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`
		author:%s
		version:%s
		commit:%s
		build time:%s
		`, Inject.Author, Inject.Version, Inject.Commit, Inject.BuildTime)
	},
}
