#!/usr/bin/env bash
set -euo pipefail

DIR=$(cd "$(dirname "$0")" && pwd)
SERVER_DIR="$DIR/server"
WEB_DIR="$DIR/web"
BIN_DIR="$SERVER_DIR/bin"

mkdir -p "$BIN_DIR"

cd "$SERVER_DIR"
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "$BIN_DIR/deedmx" ./
cp "config.yaml" "$BIN_DIR/"
cd "$DIR"

cd "$WEB_DIR"
if [ -f package.json ]; then
  if command -v pnpm >/dev/null 2>&1; then pnpm build; elif command -v npm >/dev/null 2>&1; then npm run build; elif command -v yarn >/dev/null 2>&1; then yarn build; fi
fi
cd "$DIR"

REMOTE_SERVER="${REMOTE_SERVER:-110.41.59.11}"
REMOTE_USER="${REMOTE_USER:-root}"
REMOTE_PATH="${REMOTE_PATH:-/data/deedmx/}"
REMOTE_DIST_PATH="${REMOTE_DIST_PATH:-/data/deedmx/dist/}"

ssh "$REMOTE_USER@$REMOTE_SERVER" "systemctl stop deedmx"
scp "$BIN_DIR/deedmx" "$REMOTE_USER@$REMOTE_SERVER:$REMOTE_PATH"
scp "$BIN_DIR/config.yaml" "$REMOTE_USER@$REMOTE_SERVER:$REMOTE_PATH"
ssh "$REMOTE_USER@$REMOTE_SERVER" "rm -rf $REMOTE_DIST_PATH/*"
scp -r "$WEB_DIR/dist" "$REMOTE_USER@$REMOTE_SERVER:$REMOTE_PATH"
ssh "$REMOTE_USER@$REMOTE_SERVER" "systemctl start deedmx"

echo success