package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Kiran-B/git-who/internal/profile"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <profile>",
	Short: "Delete a profile",
	Args:  cobra.ExactArgs(1),
	RunE:  runDelete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func runDelete(cmd *cobra.Command, args []string) error {
	name := args[0]

	store, err := profile.Load()
	if err != nil {
		color.Red("Error: %s", err)
		return nil
	}

	if store.FindByName(name) == nil {
		color.Red("Error: profile %q not found. Run \"git who list\" to see available profiles.", name)
		return nil
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Delete profile %q? (y/N): ", name)
	if !scanner.Scan() {
		return nil
	}

	answer := strings.TrimSpace(strings.ToLower(scanner.Text()))
	if answer != "y" && answer != "yes" {
		color.Yellow("Cancelled.")
		return nil
	}

	if err := profile.Delete(name); err != nil {
		color.Red("Error: %s", err)
		return nil
	}

	color.Green("Profile %q deleted.", name)
	return nil
}
