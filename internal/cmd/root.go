package cmd

import (
	"fmt"
	"strings"

	"github.com/Kiran-B/git-who/internal/gitconfig"
	"github.com/Kiran-B/git-who/internal/profile"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-who",
	Short: "Manage and switch Git identities per repository",
	RunE:  runRoot,
}

func Execute() error {
	return rootCmd.Execute()
}

func runRoot(cmd *cobra.Command, args []string) error {
	name, _ := readIdentity("user.name")
	email, _ := readIdentity("user.email")

	if name == "" && email == "" {
		color.Yellow("No Git identity configured. Run \"git who use <profile>\" to set one.")
		return nil
	}

	signingKey, _ := readIdentity("user.signingKey")
	sshCommand, _ := readIdentity("core.sshCommand")
	sshKey := extractSSHKey(sshCommand)

	store, _ := profile.Load()
	var matched *profile.Profile
	for i := range store.Profiles {
		p := &store.Profiles[i]
		if p.FullName == name && p.Email == email {
			matched = p
			break
		}
	}

	var parts []string
	parts = append(parts, fmt.Sprintf("%s <%s>", name, email))

	if matched != nil {
		parts = append(parts, fmt.Sprintf(" [%s]", matched.Name))
	} else {
		parts = append(parts, "  (no matching profile)")
	}

	if signingKey != "" {
		parts = append(parts, fmt.Sprintf(" GPG: %s", signingKey))
	}
	if sshKey != "" {
		parts = append(parts, fmt.Sprintf(" SSH: %s", sshKey))
	}

	fmt.Println(strings.Join(parts, ""))
	return nil
}

func readIdentity(key string) (string, error) {
	val, err := gitconfig.ReadLocal(key)
	if err != nil || val != "" {
		return val, err
	}
	return gitconfig.ReadGlobal(key)
}

func extractSSHKey(sshCommand string) string {
	parts := strings.Fields(sshCommand)
	for i, p := range parts {
		if p == "-i" && i+1 < len(parts) {
			return parts[i+1]
		}
	}
	return ""
}
