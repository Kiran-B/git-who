# git-who

> Quickly switch Git identities (name, email, SSH key, GPG key) per repo

`git-who` is a lightweight macOS CLI tool for developers who maintain multiple Git identities — work, personal, open source, freelance — and need to switch between them fast, per repository.

```
$ git who
Jane Doe <jane@company.com>  [work]  GPG: ABC123  SSH: ~/.ssh/id_work
```

---

## Why

Git doesn't know who you are. You do. But remembering to set `user.name`, `user.email`, `user.signingKey`, and `core.sshCommand` every time you clone a new repo is tedious and error-prone.

`git-who` keeps your identities in one place and applies them with a single command.

---

## Installation

### Homebrew _(coming soon)_

```sh
brew install git-who
```

### Build from source

```sh
git clone https://github.com/Kiran-B/git-who.git
cd git-who
make build
make install   # copies binary to /usr/local/bin
```

Requires Go 1.23+.

---

## Usage

### Show current identity

```sh
git who
```

Shows the effective identity for the current repo (local config takes precedence over global).

### List all profiles

```sh
git who list
```

```
  work        Jane Doe <jane@company.com>
  personal    Jane D <jane@personal.dev>
* oss         janed <jane@opensource.org>   ← active in this repo
```

### Apply a profile to the current repo

```sh
git who use work
```

Sets `user.name`, `user.email`, `user.signingKey`, and `core.sshCommand` in the local `.git/config`. Does not touch your global git config.

### Add a new profile

```sh
git who add
```

Interactive prompt:

```
Profile name:   freelance
Full name:      Jane Doe
Email:          jane@clientwork.com
SSH key path:   ~/.ssh/id_freelance    (optional)
GPG key ID:     DEF456                 (optional)
```

### Edit a profile

```sh
git who edit freelance
```

### Delete a profile

```sh
git who delete freelance
```

---

## Profiles

Profiles are stored in `~/.config/git-who/profiles.json`:

```json
{
  "profiles": [
    {
      "name": "work",
      "full_name": "Jane Doe",
      "email": "jane@company.com",
      "ssh_key": "~/.ssh/id_work",
      "gpg_key": "ABC123"
    },
    {
      "name": "personal",
      "full_name": "Jane D",
      "email": "jane@personal.dev",
      "ssh_key": "~/.ssh/id_personal",
      "gpg_key": ""
    }
  ]
}
```

You can edit this file directly or use `git who add` / `git who edit`.

---

## What `git who use` does under the hood

```sh
git config user.name      "Jane Doe"
git config user.email     "jane@company.com"
git config user.signingKey "ABC123"
git config core.sshCommand "ssh -i ~/.ssh/id_work -F /dev/null"
```

All changes are **local to the current repository**. Your global `~/.gitconfig` is never modified.

---

## Roadmap

- [ ] `git who use <profile> --global` — apply profile globally
- [ ] `git who current` — machine-readable output (JSON)
- [ ] Shell completions (zsh, bash, fish)
- [ ] Homebrew tap
- [ ] Auto-apply profile based on repo remote URL pattern
