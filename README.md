# dotfiles

Declarative macOS setup with [Determinate Nix](https://determinate.systems/nix/),
[nix-darwin](https://github.com/nix-darwin/nix-darwin), and
[home-manager](https://github.com/nix-community/home-manager), plus Homebrew
(managed declaratively via [nix-homebrew](https://github.com/zhaofengli/nix-homebrew))
for GUI apps and macOS-only tools.

## What's in the box

| File | Purpose |
|------|---------|
| `flake.nix` | Entry point — wires nix-darwin, home-manager, and nix-homebrew together. Username is set once here (`user = "ramakant"`). |
| `configuration.nix` | System level: macOS defaults (Dock, Finder, keyboard), Homebrew casks/brews. |
| `home.nix` | User level: CLI packages, zsh, starship prompt, aliases, config symlinks. |
| `home/.config/` | Real config files (nvim, ghostty) — symlinked into `~/.config` so edits apply in place, no rebuild needed. |
| `rebuild.sh` | One command to apply everything. |
| `PACKAGES.md` | Menu of packages previously installed on this machine, for reference. |

Key tools installed: neovim, git, gh, lazygit, ripgrep/fd/fzf, tmux, zoxide,
eza/bat, uv, Node.js 22, docker CLI + compose, podman (CLI + Podman Desktop),
Claude Code, Codex, Ghostty, Raycast, JetBrains Mono Nerd Font.

## Supported macOS versions

Works on **macOS Sequoia (15) and newer** (verified on macOS 26), Apple Silicon
by default. For an Intel Mac, change one line in `configuration.nix`:

```nix
nixpkgs.hostPlatform = "x86_64-darwin"; # instead of aarch64-darwin
```

The Determinate Nix installer handles the Sequoia-specific build-user UID
changes automatically — just use a current installer.

## Fresh install

1. **Install Xcode command line tools** (gives you `git`):

   ```sh
   xcode-select --install
   ```

2. **Install Determinate Nix** (the config assumes it — `nix.enable = false`
   lets Determinate manage the daemon):

   ```sh
   curl -fsSL https://install.determinate.systems/nix | sh -s -- install --determinate
   ```

   Open a new terminal afterwards so `nix` is on your PATH.

3. **Homebrew**: not required beforehand. `nix-homebrew` installs/manages
   `/opt/homebrew` itself, and `autoMigrate = true` adopts an existing install
   without losing packages.

4. **Clone and rebuild**:

   ```sh
   git clone https://github.com/<you>/dotfiles.git ~/github/dotfiles
   cd ~/github/dotfiles
   ./rebuild.sh
   ```

   The script symlinks the repo to `~/.dotfiles`, bootstraps nix-darwin on the
   first run (`nix run nix-darwin -- switch`), and uses the installed
   `darwin-rebuild` from then on. It needs `sudo` (it will prompt).

5. **New username or hostname?** Edit `user` in `flake.nix`. The flake target
   is `#mac`, independent of your actual hostname.

6. Log out/in (or reboot) once so macOS defaults (Dock, keyboard, menu bar)
   fully apply.

## Day-to-day: applying changes

Edit `configuration.nix` / `home.nix`, then:

```sh
./rebuild.sh
```

- Add a GUI app → `homebrew.casks` in `configuration.nix`
- Add a CLI tool → `home.packages` in `home.nix` (prefer nixpkgs), or
  `homebrew.brews` if it's macOS-specific
- nvim/ghostty config → edit files under `home/.config/` directly; they're
  symlinked, no rebuild needed

⚠️ `homebrew.onActivation.cleanup = "zap"` **uninstalls any brew package not
declared in `configuration.nix`**. Never `brew install` manually — declare it,
or it disappears on the next rebuild.

## Reinstall / recover on an existing machine

Same as a fresh install — everything is idempotent:

```sh
git clone <repo> && cd dotfiles && ./rebuild.sh
```

- Existing dotfiles that would be clobbered (e.g. an old `~/.config/ghostty`)
  are backed up as `<name>.backup` (`backupFileExtension`).
- An existing Homebrew install is adopted, not wiped (`autoMigrate = true`) —
  but remember the zap rule above: undeclared brew packages get removed.
- To roll back a bad rebuild:

  ```sh
  sudo /run/current-system/sw/bin/darwin-rebuild --rollback
  ```

- To update all inputs (nixpkgs, home-manager, …):

  ```sh
  nix flake update && ./rebuild.sh
  ```

## Troubleshooting

- **`darwin-rebuild: command not found`** — first build hasn't succeeded yet;
  `rebuild.sh` falls back to `nix run nix-darwin` automatically, just run it.
- **Homebrew cask conflicts** (`already exists`) — activation runs with
  `--force`, so pre-existing manually-installed apps are taken over.
- **Podman**: after the first install run `podman machine init && podman machine start`
  once. The `docker` CLI talks to podman when Podman Desktop's
  "Docker Compatibility" is enabled.
