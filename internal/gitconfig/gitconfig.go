package gitconfig

import (
	"os/exec"
	"strings"
)

func ReadLocal(key string) (string, error) {
	return readConfig("--local", key)
}

func ReadGlobal(key string) (string, error) {
	return readConfig("--global", key)
}

func readConfig(scope, key string) (string, error) {
	out, err := exec.Command("git", "config", scope, key).Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 1 {
			return "", nil
		}
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func SetLocal(key, value string) error {
	return exec.Command("git", "config", "--local", key, value).Run()
}

func UnsetLocal(key string) error {
	err := exec.Command("git", "config", "--local", "--unset", key).Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 5 {
			return nil
		}
	}
	return err
}

func IsInsideGitRepo() bool {
	err := exec.Command("git", "rev-parse", "--is-inside-work-tree").Run()
	return err == nil
}
