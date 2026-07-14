#!/usr/bin/env bash
set -euo pipefail

PVYAI_REPO="${PVYAI_REPO:-pvyswiss/pvyai-coding-agent}"
PVYAI_VERSION="${PVYAI_VERSION:-latest}"
PVYAI_INSTALL_DIR="${PVYAI_INSTALL_DIR:-$HOME/.local/bin}"
PVYAI_GITHUB_API="${PVYAI_GITHUB_API:-https://api.github.com}"
PVYAI_GITHUB_BASE_URL="${PVYAI_GITHUB_BASE_URL:-https://github.com}"

usage() {
  cat <<'EOF'
Install PVYai from GitHub Releases.

Usage:
  scripts/install.sh [--version <version>] [--repo <owner/repo>] [--install-dir <path>]

Environment:
  PVYAI_VERSION          Release version or tag. Defaults to latest.
  PVYAI_REPO             GitHub repository. Defaults to pvyswiss/pvyai-coding-agent.
  PVYAI_INSTALL_DIR      Directory for the pvyai binary. Defaults to ~/.local/bin.
  PVYAI_GITHUB_API       GitHub API base URL. Defaults to https://api.github.com.
  PVYAI_GITHUB_BASE_URL  GitHub web base URL. Defaults to https://github.com.
EOF
}

fail() {
  echo "pvyai install: $*" >&2
  exit 1
}

while [ "$#" -gt 0 ]; do
  case "$1" in
    --version)
      [ "$#" -ge 2 ] || fail "--version requires a value"
      PVYAI_VERSION="$2"
      shift 2
      ;;
    --repo)
      [ "$#" -ge 2 ] || fail "--repo requires a value"
      PVYAI_REPO="$2"
      shift 2
      ;;
    --install-dir)
      [ "$#" -ge 2 ] || fail "--install-dir requires a value"
      PVYAI_INSTALL_DIR="$2"
      shift 2
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      fail "unknown argument: $1"
      ;;
  esac
done

need_command() {
  command -v "$1" >/dev/null 2>&1 || fail "$1 is required"
}

download() {
  local url="$1"
  local output="$2"

  if command -v curl >/dev/null 2>&1; then
    curl --fail --location --show-error --silent "$url" --output "$output"
  elif command -v wget >/dev/null 2>&1; then
    wget --quiet "$url" --output-document "$output"
  else
    fail "curl or wget is required"
  fi
}

download_json() {
  local url="$1"
  local output="$2"

  if command -v curl >/dev/null 2>&1; then
    curl --fail --location --show-error --silent --header 'Accept: application/vnd.github+json' "$url" --output "$output"
  elif command -v wget >/dev/null 2>&1; then
    wget --quiet --header='Accept: application/vnd.github+json' "$url" --output-document "$output"
  else
    fail "curl or wget is required"
  fi
}

detect_platform() {
  case "$(uname -s)" in
    Linux) echo "linux" ;;
    Darwin) echo "macos" ;;
    *) fail "unsupported platform: $(uname -s)" ;;
  esac
}

detect_arch() {
  case "$(uname -m)" in
    x86_64|amd64) echo "x64" ;;
    arm64|aarch64) echo "arm64" ;;
    *) fail "unsupported architecture: $(uname -m)" ;;
  esac
}

latest_tag() {
  local metadata_file="$1"
  local api_url="${PVYAI_GITHUB_API%/}/repos/${PVYAI_REPO}/releases/latest"
  local tag

  download_json "$api_url" "$metadata_file"
  tag="$(sed -n 's/.*"tag_name"[[:space:]]*:[[:space:]]*"\([^"]*\)".*/\1/p' "$metadata_file" | head -n 1)"
  [ -n "$tag" ] || fail "could not read tag_name from $api_url"
  echo "$tag"
}

verify_checksum() {
  local checksum_file="$1"

  if command -v shasum >/dev/null 2>&1; then
    shasum -a 256 -c "$checksum_file"
  elif command -v sha256sum >/dev/null 2>&1; then
    sha256sum -c "$checksum_file"
  else
    fail "shasum or sha256sum is required"
  fi
}

find_extracted_entry() {
  local root="$1"
  local name="$2"
  local kind="$3"
  local candidate

  if [ "$kind" = "dir" ] && [ -d "$root/$name" ]; then
    echo "$root/$name"
    return 0
  fi
  if [ "$kind" = "file" ] && [ -f "$root/$name" ]; then
    echo "$root/$name"
    return 0
  fi

  for candidate in "$root"/*/"$name"; do
    if [ "$kind" = "dir" ] && [ -d "$candidate" ]; then
      echo "$candidate"
      return 0
    fi
    if [ "$kind" = "file" ] && [ -f "$candidate" ]; then
      echo "$candidate"
      return 0
    fi
  done

  return 1
}

find_extracted_binary() {
  find_extracted_entry "$1" "pvyai" "file"
}

copy_optional_file() {
  local name="$1"
  local source_path

  if source_path="$(find_extracted_entry "$extract_dir" "$name" "file")"; then
    cp "$source_path" "$PVYAI_INSTALL_DIR/$name"
    chmod 755 "$PVYAI_INSTALL_DIR/$name"
  fi
}

copy_optional_dir() {
  local name="$1"
  local source_path

  if source_path="$(find_extracted_entry "$extract_dir" "$name" "dir")"; then
    rm -rf "$PVYAI_INSTALL_DIR/$name"
    cp -R "$source_path" "$PVYAI_INSTALL_DIR/$name"
  fi
}

need_command uname
need_command sed
need_command tar
need_command mktemp

tmp_dir="$(mktemp -d "${TMPDIR:-/tmp}/pvyai-install.XXXXXX")"
cleanup() {
  rm -rf "$tmp_dir"
}
trap cleanup EXIT

if [ "$PVYAI_VERSION" = "latest" ]; then
  tag="$(latest_tag "$tmp_dir/latest.json")"
else
  case "$PVYAI_VERSION" in
    v*) tag="$PVYAI_VERSION" ;;
    *) tag="v$PVYAI_VERSION" ;;
  esac
fi

version="${tag#v}"
platform="$(detect_platform)"
arch="$(detect_arch)"
archive_name="pvyai-v${version}-${platform}-${arch}.tar.gz"
checksum_name="${archive_name}.sha256"
release_url="${PVYAI_GITHUB_BASE_URL%/}/${PVYAI_REPO}/releases/download/${tag}"
archive_path="$tmp_dir/$archive_name"
checksum_path="$tmp_dir/$checksum_name"
extract_dir="$tmp_dir/extract"

echo "Installing PVYai ${tag} for ${platform}-${arch}"
download "${release_url}/${archive_name}" "$archive_path"
download "${release_url}/${checksum_name}" "$checksum_path"

(
  cd "$tmp_dir"
  verify_checksum "$checksum_name"
)

mkdir -p "$extract_dir"
tar -xzf "$archive_path" -C "$extract_dir"

binary_path="$(find_extracted_binary "$extract_dir")" || fail "release archive did not contain pvyai"

mkdir -p "$PVYAI_INSTALL_DIR"
cp "$binary_path" "$PVYAI_INSTALL_DIR/pvyai"
chmod 755 "$PVYAI_INSTALL_DIR/pvyai"
copy_optional_file "pvyai-linux-sandbox"
copy_optional_file "pvyai-seccomp"
copy_optional_dir "helpers"

echo "Installed $PVYAI_INSTALL_DIR/pvyai"

case ":$PATH:" in
  *":$PVYAI_INSTALL_DIR:"*) ;;
  *) echo "Add $PVYAI_INSTALL_DIR to PATH to run pvyai from any directory." ;;
esac
