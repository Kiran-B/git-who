package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-who",
	Short: "Manage and switch Git identities per repository",
}

func Execute() error {
	return rootCmd.Execute()
}
