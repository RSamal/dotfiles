# Package Reference

A catalog of every Homebrew package that was installed on this machine, with a
plain-English description of what each one is and what it's used for.

Use this as a menu: when you need a tool, add its name to the matching array in
`configuration.nix` and run `./rebuild.sh`. When you no longer need it, delete
the line and rebuild.

- **Casks** = GUI applications → go in `homebrew.casks`
- **Formulae** = command-line tools / libraries → go in `homebrew.brews`

> Reminder: `homebrew.onActivation.cleanup = "zap"` means anything **not**
> listed in `configuration.nix` gets uninstalled on the next rebuild. The config
> is the single source of truth.

---

## Casks (GUI apps)

### Currently declared in `configuration.nix`

| Cask | What it is / used for |
|------|-----------------------|
| `codex` | OpenAI Codex — AI coding agent CLI/app. |
| `font-jetbrains-mono-nerd-font` | JetBrains Mono font patched with Nerd Font icons (needed by terminals, editors, prompts like Starship/Powerlevel10k). |
| `ghostty` | Fast, GPU-accelerated terminal emulator. |
| `hyperwhisper` | Voice-to-text dictation app (speech → text anywhere). |
| `openlens` | Kubernetes IDE / dashboard (open-source build of Lens). |
| `podman-desktop` | Desktop GUI for containers & Kubernetes; a Docker Desktop alternative. |
| `raycast` | Launcher and productivity tool; a powerful Spotlight replacement. |

### Removed (were installed, not kept)

| Cask | What it is / used for |
|------|-----------------------|
| `headlamp` | Kubernetes dashboard UI (overlaps with openlens). |
| `langgraph-studio` | Visual IDE for building/debugging LangGraph agent workflows. |
| `wezterm` | GPU-accelerated terminal emulator (overlaps with ghostty). |

---

## Formulae (command-line tools)

None are declared yet. Add the ones you want to a `brews = [ ... ]` block in
`configuration.nix`. Grouped by purpose below.

### General-purpose CLI (recommended on any machine)

| Formula | What it is / used for |
|---------|-----------------------|
| `gh` | GitHub CLI — manage PRs, issues, repos from the terminal. |
| `eza` | Modern replacement for `ls` (colors, icons, git status). |
| `neovim` | Modern, extensible Vim-based text editor. |
| `tmux` | Terminal multiplexer — split panes, persistent sessions. |
| `stow` | GNU Stow — symlink manager for dotfiles (relevant to this repo). |
| `git-lfs` | Git Large File Storage — version large binary files in git. |
| `git-filter-repo` | Rewrite git history (remove files, scrub secrets, split repos). |
| `tealdeer` | Fast Rust client for `tldr` — simplified command examples. |
| `navi` | Interactive cheatsheet tool for shell commands. |
| `chafa` | Render images as ASCII/ANSI art in the terminal. |
| `pastel` | Generate, analyze, and convert colors from the CLI. |
| `zsh-autosuggestions` | Fish-like inline command suggestions for zsh. |
| `zsh-completions` | Extra tab-completion definitions for zsh. |
| `zsh-syntax-highlighting` | Real-time command syntax highlighting in zsh. |

### Python toolchain

| Formula | What it is / used for |
|---------|-----------------------|
| `uv` | Extremely fast Python package/venv manager (pip/venv replacement). |
| `pipx` | Install and run Python CLI apps in isolated environments. |
| `poetry` | Python dependency management and packaging. |
| `pytest` | Python testing framework. |
| `pre-commit` | Framework for managing git pre-commit hooks (linters/formatters). |

### Build toolchains (compiling native code)

| Formula | What it is / used for |
|---------|-----------------------|
| `cmake` | Cross-platform build-system generator. |
| `make` | Classic build automation tool (Makefiles). |
| `ninja` | Small, fast build system (often used with cmake). |
| `lld` | LLVM's fast linker. |
| `go` | Go programming language toolchain. |
| `zig@0.15` | Zig language toolchain (also a C/C++ compiler), pinned to 0.15. |
| `vite` | Frontend build tool / dev server. **Needs Node.js (not in this list).** |

### Containers & Kubernetes

| Formula | What it is / used for |
|---------|-----------------------|
| `docker` | Docker CLI — build and run containers. |
| `docker-compose` | Define and run multi-container apps via YAML. |
| `kind` | Kubernetes IN Docker — run local k8s clusters. |
| `kubernetes-cli` | `kubectl` — the Kubernetes command-line client. |
| `stern` | Tail logs across multiple pods/containers at once. |
| `openshift-cli` | `oc` — CLI for Red Hat OpenShift clusters. |
| `teleport` | Secure access proxy for SSH/Kubernetes/databases. |

### Databases & services (run background daemons)

| Formula | What it is / used for |
|---------|-----------------------|
| `postgresql@14` | PostgreSQL relational database, pinned to v14. |
| `redis` | In-memory key-value data store / cache. |
| `nginx` | Web server / reverse proxy. |
| `mpd` | Music Player Daemon — server-side music playback. |
| `spark` | Apache Spark — distributed data processing engine. |

### Niche / uncategorized (evaluate before adding)

| Formula | What it is / used for |
|---------|-----------------------|
| `cppman` | C++ manual pages (cppreference) in `man` format. |
| `langgraph-cli` | CLI for LangGraph (paired with the removed langgraph-studio app). |
| `stormy` | Small terminal weather app. |
| `taproom` | Homebrew tap/formula helper utility. |

---

## Taps (third-party package repositories)

Taps that were configured. A tap only matters if you install a package *from* it;
declare taps in a `taps = [ ... ]` block only when you need one of their packages.

| Tap | Provides (examples) |
|-----|---------------------|
| `homebrew/services` | `brew services` — manage background daemons (postgres, redis, etc.). |
| `oven-sh/bun` | `bun` — fast JavaScript runtime & package manager. |
| `derailed/k9s` | `k9s` — terminal UI for Kubernetes clusters. |
| `sass/sass` | `sass` — CSS preprocessor. |
| `dart-lang/dart` | `dart` — Dart language SDK. |
| `koekeishiya/formulae` | `yabai` / `skhd` — tiling window manager & hotkey daemon. |
| `louisbrunner/valgrind` | `valgrind` — memory debugging/profiling (macOS build). |
| `asmvik/formulae` | Custom formulae from this maintainer. |

---

## How to add / remove a package

1. Open `configuration.nix`.
2. Add the name to the right array:
   - GUI app → `homebrew.casks`
   - CLI tool → `homebrew.brews`
   - Repository → `homebrew.taps`
3. Run `./rebuild.sh`.
4. To remove: delete the line and rebuild — `cleanup = "zap"` uninstalls it.
