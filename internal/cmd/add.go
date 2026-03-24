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

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new profile",
	RunE:  runAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAdd(cmd *cobra.Command, args []string) error {
	scanner := bufio.NewScanner(os.Stdin)

	name := promptRequired(scanner, "Profile name")
	if name == "" {
		return nil
	}

	store, err := profile.Load()
	if err != nil {
		color.Red("Error: %s", err)
		return nil
	}
	if store.FindByName(name) != nil {
		color.Red("Error: a profile named %q already exists", name)
		return nil
	}

	fullName := promptRequired(scanner, "Full name")
	if fullName == "" {
		return nil
	}

	email := promptRequired(scanner, "Email")
	if email == "" {
		return nil
	}

	sshKey := promptOptional(scanner, "SSH key path (optional)")
	gpgKey := promptOptional(scanner, "GPG key ID (optional)")

	p := profile.Profile{
		Name:     name,
		FullName: fullName,
		Email:    email,
		SSHKey:   sshKey,
		GPGKey:   gpgKey,
	}

	if err := profile.Add(p); err != nil {
		color.Red("Error: %s", err)
		return nil
	}

	color.Green("Profile %q created.", name)
	return nil
}

func promptRequired(scanner *bufio.Scanner, label string) string {
	for {
		fmt.Printf("%s: ", label)
		if !scanner.Scan() {
			return ""
		}
		val := strings.TrimSpace(scanner.Text())
		if val != "" {
			return val
		}
		color.Yellow("%s is required.", label)
	}
}

func promptOptional(scanner *bufio.Scanner, label string) string {
	fmt.Printf("%s: ", label)
	if !scanner.Scan() {
		return ""
	}
	return strings.TrimSpace(scanner.Text())
}
