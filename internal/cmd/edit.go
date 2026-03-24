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

var editCmd = &cobra.Command{
	Use:   "edit <profile>",
	Short: "Edit an existing profile",
	Args:  cobra.ExactArgs(1),
	RunE:  runEdit,
}

func init() {
	rootCmd.AddCommand(editCmd)
}

func runEdit(cmd *cobra.Command, args []string) error {
	name := args[0]

	store, err := profile.Load()
	if err != nil {
		color.Red("Error: %s", err)
		return nil
	}

	p := store.FindByName(name)
	if p == nil {
		color.Red("Error: profile %q not found. Run \"git who list\" to see available profiles.", name)
		return nil
	}

	scanner := bufio.NewScanner(os.Stdin)

	fullName := promptWithDefault(scanner, "Full name", p.FullName)
	email := promptWithDefault(scanner, "Email", p.Email)
	sshKey := promptWithDefault(scanner, "SSH key path", p.SSHKey)
	gpgKey := promptWithDefault(scanner, "GPG key ID", p.GPGKey)

	updated := profile.Profile{
		Name:     p.Name,
		FullName: fullName,
		Email:    email,
		SSHKey:   sshKey,
		GPGKey:   gpgKey,
	}

	if err := profile.Update(name, updated); err != nil {
		color.Red("Error: %s", err)
		return nil
	}

	color.Green("Profile %q updated.", name)
	return nil
}

func promptWithDefault(scanner *bufio.Scanner, label, defaultVal string) string {
	if defaultVal != "" {
		fmt.Printf("%s [%s]: ", label, defaultVal)
	} else {
		fmt.Printf("%s: ", label)
	}
	if !scanner.Scan() {
		return defaultVal
	}
	val := strings.TrimSpace(scanner.Text())
	if val == "" {
		return defaultVal
	}
	return val
}
