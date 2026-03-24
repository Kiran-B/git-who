package cmd

import (
	"fmt"

	"github.com/Kiran-B/git-who/internal/gitconfig"
	"github.com/Kiran-B/git-who/internal/profile"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all saved profiles",
	RunE:  runList,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func runList(cmd *cobra.Command, args []string) error {
	store, err := profile.Load()
	if err != nil {
		color.Red("Error: %s", err)
		return nil
	}

	if len(store.Profiles) == 0 {
		color.Yellow("No profiles saved. Run \"git who add\" to create one.")
		return nil
	}

	var currentName, currentEmail string
	if gitconfig.IsInsideGitRepo() {
		currentName, _ = readIdentity("user.name")
		currentEmail, _ = readIdentity("user.email")
	}

	for _, p := range store.Profiles {
		active := p.FullName == currentName && p.Email == currentEmail
		marker := " "
		if active {
			marker = "*"
		}
		fmt.Printf("%s %-12s %s <%s>\n", marker, p.Name, p.FullName, p.Email)
	}

	return nil
}
