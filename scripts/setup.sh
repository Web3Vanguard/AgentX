#!/usr/bin/env bash

set -euo pipefail

# Remote-friendly installer for AgentX
# Usage:
#  curl -fsSL https://raw.githubusercontent.com/Web3Vanguard/AgentX/tree/main/scripts/install.sh | bash
# or
#  curl -fsSL https://raw.githubusercontent.com/Web3Vanguard/AgentX/tree/main/scripts/install.sh | sudo bash

BIN_NAME="agentX"
RELEASE_TAG="v0.0.1"  # release tag you provided
DEFAULT_INSTALL_DIR="/usr/local/bin"
FALLBACK_INSTALL_DIR="$HOME/.local/bin"

progress_bar() {
  local progress=$1
  local total=100
  local width=40
  local filled=$(( progress * width / total ))
  local empty=$(( width - filled ))
  printf "\r["
  printf "%0.s#" $(seq 1 $filled)
  printf "%0.s-" $(seq 1 $empty)
  printf "] %d%%" "$progress"
}

run_with_progress() {
  local duration=$1
  local message=$2
  echo -ne "\n$message"
  # fractional sleep requires bc; fall back to plain sleep if not present
  if command -v bc >/dev/null 2>&1; then
    for i in $(seq 1 100); do
      progress_bar "$i"
      sleep "$(bc -l <<< "$duration/100")"
    done
  else
    # coarse progress if bc not available
    local step_sleep=$(awk "BEGIN {printf \"%.3f\", $duration/10}")
    for i in $(seq 1 10); do
      progress_bar $(( i * 10 ))
      sleep "$step_sleep"
    done
  fi
  echo ""
}

# Detect OS and ARCH
OS_RAW="$(uname -s)"
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH_RAW="$(uname -m)"

case "$ARCH_RAW" in
  x86_64|amd64) ARCH="amd64" ;;        # matches your asset naming (amd64)
  aarch64|arm64) ARCH="arm64" ;;
  armv7l) ARCH="armv7" ;;
  *) ARCH="$ARCH_RAW" ;;
esac

# Map OS to asset name: your assets use "darwin" for macOS
case "$OS" in
  linux*)   ASSET_OS="linux" ;;
  darwin*)  ASSET_OS="darwin" ;;
  mingw*|msys*|cygwin*) ASSET_OS="windows" ;;
  *) echo "Unsupported OS: $OS_RAW" >&2; exit 1 ;;
esac

# Build download URL for the specific tag (not `latest`)
BASE_URL="${SOMNIA_BASE_URL:-https://github.com/Web3Vanguard/AgentX/releases/download/${RELEASE_TAG}}"
ASSET_NAME="${BIN_NAME}-${ASSET_OS}-${ARCH}"
# Windows assets have .exe suffix
if [[ "$ASSET_OS" == "windows" ]]; then
  ASSET_NAME="${ASSET_NAME}.exe"
fi
DOWNLOAD_URL="_
