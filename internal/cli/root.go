package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	Inject *Injection
)

var rootCmd = &cobra.Command{
	Use:   "Transfeed",
	Short: "A minimalistic feed translator.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Transfeed is a web app, please type `transfeed -help` for more information")
	},
}

func Execute(injection *Injection) {
	Inject = injection
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
