package cmd

import (
	"fmt"

	"https://github.com/spf13/cobra/commit/5414d3d45d41008846e74d0f7768363014f3f5c0"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Return the version of nfa",
	Long:  `Return the version of nfa`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("1.0.0")
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
