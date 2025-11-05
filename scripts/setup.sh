#!/usr/bin/env bash

set -euo pipefail

# Remote-friendly installer for AgentX
# Usage:
#  curl -fsSL https://raw.githubusercontent.com/Web3Vanguard/AgentX/main/scripts/install.sh | bash
# or
#  curl -fsSL https://raw.githubusercontent.com/Web3Vanguard/AgentX/main/scripts/install.sh | sudo bash

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
DOWNLOAD_URL="${BASE_URL}/${ASSET_NAME}"

echo "🔍 Installing ${BIN_NAME} for ${ASSET_OS}/${ARCH}"
echo "📦 Source: ${DOWNLOAD_URL}"

TMP_DIR="$(mktemp -d)"
cleanup() { rm -rf "$TMP_DIR"; }
trap cleanup EXIT

TARGET_TMP="$TMP_DIR/$BIN_NAME"

echo "⬇️  Downloading binary..."
# prefer curl, fallback to wget
if command -v curl >/dev/null 2>&1; then
  if ! curl -fsSL "$DOWNLOAD_URL" -o "$TARGET_TMP"; then
    echo "Download failed from: $DOWNLOAD_URL" >&2
    exit 1
  fi
elif command -v wget >/dev/null 2>&1; then
  if ! wget -qO "$TARGET_TMP" "$DOWNLOAD_URL"; then
    echo "Download failed from: $DOWNLOAD_URL" >&2
    exit 1
  fi
else
  echo "Error: neither curl nor wget is available." >&2
  exit 1
fi

# Ensure we have an executable name on disk
if [[ ! -s "$TARGET_TMP" ]]; then
  echo "Download produced an empty file. Aborting." >&2
  exit 1
fi

chmod +x "$TARGET_TMP" || true

install_with() {
  local dest_dir="$1"
  echo "📂 Installing to $dest_dir..."
  if command -v install >/dev/null 2>&1; then
    install -m 0755 "$TARGET_TMP" "$dest_dir/$BIN_NAME"
  else
    cp "$TARGET_TMP" "$dest_dir/$BIN_NAME"
    chmod 0755 "$dest_dir/$BIN_NAME"
  fi
}

echo "⚙️  Setting up binary..."
run_with_progress 1.5 "Installing..."

INSTALL_DIR="$DEFAULT_INSTALL_DIR"
if [[ -w "$INSTALL_DIR" ]]; then
  install_with "$INSTALL_DIR"
else
  if command -v sudo >/dev/null 2>&1; then
    echo "🔐 Using sudo for install to $INSTALL_DIR..."
    run_with_progress 1 "Authorizing..."
    if command -v install >/dev/null 2>&1; then
      sudo install -m 0755 "$TARGET_TMP" "$INSTALL_DIR/$BIN_NAME"
    else
      sudo cp "$TARGET_TMP" "$INSTALL_DIR/$BIN_NAME"
      sudo chmod 0755 "$INSTALL_DIR/$BIN_NAME"
    fi
  else
    echo "⚠️  No write access to $INSTALL_DIR and sudo not available; falling back to user dir: $FALLBACK_INSTALL_DIR"
    mkdir -p "$FALLBACK_INSTALL_DIR"
    install_with "$FALLBACK_INSTALL_DIR"
    case ":$PATH:" in
      *:"$FALLBACK_INSTALL_DIR":*) ;;
      *) echo "💡 Add to PATH: export PATH=\"$FALLBACK_INSTALL_DIR:\$PATH\"";;
    esac
  fi
fi

# Decide where the installed binary actually is
INSTALLED_PATH="$DEFAULT_INSTALL_DIR/$BIN_NAME"
if [[ ! -x "$INSTALLED_PATH" ]]; then
  INSTALLED_PATH="$FALLBACK_INSTALL_DIR/$BIN_NAME"
fi

# macOS quarantine removal (best-effort)
if [[ "$ASSET_OS" == "darwin" ]] && command -v xattr >/dev/null 2>&1; then
  xattr -d com.apple.quarantine "$INSTALLED_PATH" 2>/dev/null || true
fi

echo "🔍 Verifying installation..."
run_with_progress 1 "Checking binary..."

if [[ -x "$INSTALLED_PATH" ]]; then
  # Try --version, but fall back to running the binary without args if that fails
  if "$INSTALLED_PATH" --version >/dev/null 2>&1; then
    echo -e "\n✅ Installed successfully!"
    echo "📍 Location: $INSTALLED_PATH"
    echo -n "🧩 Version: "; "$INSTALLED_PATH" --version || true
  else
    echo -e "\n✅ Installed (but '$BIN_NAME --version' failed or is unsupported)."
    echo "📍 Location: $INSTALLED_PATH"
    echo "Try running directly: $INSTALLED_PATH"
  fi
else
  echo "❌ Installation failed: $INSTALLED_PATH not executable." >&2
  exit 1
fi

echo "🎉 Done!"
exit 0
