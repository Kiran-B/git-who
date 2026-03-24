package cmd

import (
	"fmt"

	"github.com/Kiran-B/git-who/internal/gitconfig"
	"github.com/Kiran-B/git-who/internal/profile"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use <profile>",
	Short: "Apply a profile to the current repository",
	Args:  cobra.ExactArgs(1),
	RunE:  runUse,
}

func init() {
	rootCmd.AddCommand(useCmd)
}

func runUse(cmd *cobra.Command, args []string) error {
	if !gitconfig.IsInsideGitRepo() {
		color.Red("Error: not inside a git repository")
		return nil
	}

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

	if err := gitconfig.SetLocal("user.name", p.FullName); err != nil {
		color.Red("Error: could not set user.name: %s", err)
		return nil
	}
	if err := gitconfig.SetLocal("user.email", p.Email); err != nil {
		color.Red("Error: could not set user.email: %s", err)
		return nil
	}

	if p.GPGKey != "" {
		if err := gitconfig.SetLocal("user.signingKey", p.GPGKey); err != nil {
			color.Red("Error: could not set user.signingKey: %s", err)
			return nil
		}
	} else {
		_ = gitconfig.UnsetLocal("user.signingKey")
	}

	if p.SSHKey != "" {
		sshCmd := fmt.Sprintf("ssh -i %s -F /dev/null", p.SSHKey)
		if err := gitconfig.SetLocal("core.sshCommand", sshCmd); err != nil {
			color.Red("Error: could not set core.sshCommand: %s", err)
			return nil
		}
	} else {
		_ = gitconfig.UnsetLocal("core.sshCommand")
	}

	color.Green("Switched to profile %q — %s <%s>", p.Name, p.FullName, p.Email)
	return nil
}
