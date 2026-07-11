#!/usr/bin/env bash
set -euo pipefail
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd -P)"
ln -sfn "$DIR" ~/.dotfiles
# darwin-rebuild only exists after the first successful switch. Bootstrap with
# `nix run` until then; afterwards use the installed binary by full path so
# sudo's secure_path doesn't hide it.
if [ -x /run/current-system/sw/bin/darwin-rebuild ]; then
  exec sudo /run/current-system/sw/bin/darwin-rebuild switch --flake ~/.dotfiles#mac
else
  exec sudo nix run nix-darwin -- switch --flake ~/.dotfiles#mac
fi
