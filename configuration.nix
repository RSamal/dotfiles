{ user, ... }:

{
  # Determinate already manages the Nix daemon, so nix-darwin shouldn't.
  nix.enable = false;

  nixpkgs.config.allowUnfree = true;
  nixpkgs.hostPlatform = "aarch64-darwin"; # use x86_64-darwin for Intel CPU

  system.primaryUser = user;
  users.users.${user} = {
    home = "/Users/${user}";
  };
  system.stateVersion = 6;
  system.defaults = {
    NSGlobalDomain = {
      AppleInterfaceStyle = "Dark";
      KeyRepeat = 2;          # fast key repeat
      InitialKeyRepeat = 15;  # short delay before repeat
      _HIHideMenuBar = true;  # auto-hide the menu bar
      AppleShowAllExtensions = true;
    };
    dock.autohide = true;
    finder.FXPreferredViewStyle = "Nlsv";  # list view by default
    finder.CreateDesktop = false;          # clean desktop
    trackpad.Clicking = true;              # tap to click
  };

  nix-homebrew = {
    enable = true;
    inherit user;
    # Take over the existing /opt/homebrew installation without losing
    # already-installed packages.
    autoMigrate = true;
  };

   homebrew = {
    enable = true;
    onActivation.cleanup = "zap";  # remove anything not declared below
    onActivation.autoUpdate = true;
    onActivation.extraFlags = [ "--force" ];
    taps = [
      # Toe tiling WM — tap lives at a custom URL (repo isn't homebrew-toe)
      { name = "theclifmeister/toe"; clone_target = "https://github.com/theclifmeister/toe"; }
    ];
    brews = [
      "podman"                         # container engine CLI (podman machine on macOS)
    ];
    casks = [
      "claude-code"                    # Anthropic Claude Code coding agent CLI
      "codex"                          # OpenAI Codex coding agent CLI
      "finicky"                        # URL router: opens links from Slack etc. in a new Chrome window
      "ghostty"                        # terminal emulator
      "theclifmeister/toe/toe"         # Toe — native Hyprland-dwindle tiling WM
      # "libreoffice"                  # app stays installed; cask definition is
      # broken with Homebrew 6 ("command_wrapper") — re-enable when fixed upstream
      "podman-desktop"                 # container/Kubernetes desktop
      "raycast"                        # launcher / Spotlight replacement
    ];
  };
}
