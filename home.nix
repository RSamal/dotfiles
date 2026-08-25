{ config, pkgs, user, ... }:

let
  dotfiles = "${config.home.homeDirectory}/.dotfiles";

  # git worktree manager (github.com/satococoa/wtp) — not in nixpkgs, so built here
  wtp = pkgs.buildGoModule rec {
    pname = "wtp";
    version = "2.10.3";
    src = pkgs.fetchFromGitHub {
      owner = "satococoa";
      repo = "wtp";
      rev = "v${version}";
      hash = "sha256-KgayKjH4iHi7LgWwk2Laba33bMVZdbiMQgSmqBSTfZ0=";
    };
    vendorHash = "sha256-zsSNo1MQgpvH3ZSd3kmvdIpOCVJgSu1/pYLltx/9dZg=";
    subPackages = [ "cmd/wtp" ];
    doCheck = false; # integration tests expect to run inside a git repo
    ldflags = [ "-s" "-w" "-X main.version=${version}" ];
  };
in

{
  home.username = user;
  home.homeDirectory = "/Users/${user}";
  home.stateVersion = "26.05";

  home.packages = with pkgs; [
    # cli i use constantly
    ripgrep   # fast search
    fd        # fast find
    fzf       # fuzzy finder
    jq        # json on the command line
    lazygit   # git TUI
    neovim     # editor
    git        # pin git via nix, not Apple's
    gh         # GitHub CLI
    eza        # modern ls
    bat        # cat with syntax highlighting
    tmux       # terminal multiplexer
    zoxide     # smarter cd
    tealdeer   # tldr — quick command examples
    uv         # fast Python package/venv manager
    nodejs_22  # Node.js 22 LTS + npm (web frontend: apps/web / vite)

    # containers — CLIs only; the engine is podman-desktop (enable its
    # "Docker Compatibility" so `docker` talks to the podman socket).
    docker-client   # `docker` CLI (client only)
    docker-compose  # `docker compose`

    # the font everything renders in
    nerd-fonts.jetbrains-mono
    yazi

    wtp  # git worktree manager (defined above)
  ];
  fonts.fontconfig.enable = true;
  home.sessionVariables.EDITOR = "nvim";

  programs.zsh = {
    enable = true;
    autosuggestion.enable = true;      # ghost text from history
    syntaxHighlighting.enable = true;  # commands turn green when valid
    initContent = ''
      bindkey '^f' autosuggest-accept
      eval "$(wtp shell-init zsh)"  # wtp cd + completions
    '';
    shellAliases = {
      ".." = "cd ..";
      add = "git add .";
      push = "git push";
      pull = "git pull";
      m = "git switch main";
      cc = "claude --dangerously-skip-permissions";
      co = "codex --full-auto";
    };
  };


   programs.starship = {
    enable = true;
    settings = {
      add_newline = false;
      format = "$directory$git_branch$git_status$cmd_duration$line_break$character";
      character = {
        success_symbol = "[❯](purple)";
        error_symbol = "[❯](red)";
      };
      cmd_duration.format = "[$duration]($style) ";
    };
  };

  # Edit-in-place: the real file stays in my repo, ~/.config just points at it.
  home.file.".config/ghostty".source =
    config.lib.file.mkOutOfStoreSymlink "${dotfiles}/home/.config/ghostty";

  home.file.".config/nvim".source =
    config.lib.file.mkOutOfStoreSymlink "${dotfiles}/home/.config/nvim";

}
