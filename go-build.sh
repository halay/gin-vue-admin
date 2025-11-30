#!/usr/bin/env bash
set -euo pipefail

DIR=$(cd "$(dirname "$0")" && pwd)
SERVER_DIR="$DIR/server"
BIN_DIR="$SERVER_DIR/bin"

mkdir -p "$BIN_DIR"

cd "$SERVER_DIR"
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "$BIN_DIR/deedmx" ./
cp "config.yaml" "$BIN_DIR/"
cd "$DIR"


REMOTE_SERVER="${REMOTE_SERVER:-110.41.59.11}"
REMOTE_USER="${REMOTE_USER:-root}"
REMOTE_PATH="${REMOTE_PATH:-/data/deedmx/}"

ssh "$REMOTE_USER@$REMOTE_SERVER" "systemctl stop deedmx"
scp "$BIN_DIR/deedmx" "$REMOTE_USER@$REMOTE_SERVER:$REMOTE_PATH"
scp "$BIN_DIR/config.yaml" "$REMOTE_USER@$REMOTE_SERVER:$REMOTE_PATH"
ssh "$REMOTE_USER@$REMOTE_SERVER" "systemctl start deedmx"

echo success