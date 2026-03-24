package profile

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Profile struct {
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	SSHKey   string `json:"ssh_key"`
	GPGKey   string `json:"gpg_key"`
}

type ProfileStore struct {
	Profiles []Profile `json:"profiles"`
}

func ConfigDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not determine home directory: %w", err)
	}
	return filepath.Join(home, ".config", "git-who"), nil
}

func ConfigPath() (string, error) {
	dir, err := ConfigDir()
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", fmt.Errorf("could not create config directory: %w", err)
	}
	return filepath.Join(dir, "profiles.json"), nil
}

func Load() (ProfileStore, error) {
	path, err := ConfigPath()
	if err != nil {
		return ProfileStore{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return ProfileStore{}, nil
		}
		return ProfileStore{}, fmt.Errorf("could not read profiles file — try checking %s", path)
	}

	var store ProfileStore
	if err := json.Unmarshal(data, &store); err != nil {
		return ProfileStore{}, fmt.Errorf("could not read profiles file — try checking %s", path)
	}
	return store, nil
}

func Save(store ProfileStore) error {
	path, err := ConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(store, "", "  ")
	if err != nil {
		return fmt.Errorf("could not save profiles: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("could not write profiles file: %w", err)
	}
	return nil
}

func (s *ProfileStore) FindByName(name string) *Profile {
	for i := range s.Profiles {
		if strings.EqualFold(s.Profiles[i].Name, name) {
			return &s.Profiles[i]
		}
	}
	return nil
}

func Add(p Profile) error {
	p.Name = strings.TrimSpace(p.Name)

	store, err := Load()
	if err != nil {
		return err
	}

	if store.FindByName(p.Name) != nil {
		return fmt.Errorf("a profile named %q already exists", p.Name)
	}

	store.Profiles = append(store.Profiles, p)
	return Save(store)
}

func Update(name string, updated Profile) error {
	store, err := Load()
	if err != nil {
		return err
	}

	for i := range store.Profiles {
		if strings.EqualFold(store.Profiles[i].Name, name) {
			store.Profiles[i] = updated
			return Save(store)
		}
	}
	return fmt.Errorf("profile %q not found", name)
}

func Delete(name string) error {
	store, err := Load()
	if err != nil {
		return err
	}

	found := false
	filtered := make([]Profile, 0, len(store.Profiles))
	for _, p := range store.Profiles {
		if strings.EqualFold(p.Name, name) {
			found = true
			continue
		}
		filtered = append(filtered, p)
	}

	if !found {
		return fmt.Errorf("profile %q not found", name)
	}

	store.Profiles = filtered
	return Save(store)
}
