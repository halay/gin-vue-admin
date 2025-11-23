DIR=$(cd "$(dirname "$0")" && pwd)
SERVER_DIR="$DIR/server"
WEB_DIR="$DIR/web"
BIN_DIR="$SERVER_DIR/bin"

mkdir -p "$BIN_DIR"

cd "$SERVER_DIR"
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "$BIN_DIR/deedmx" ./