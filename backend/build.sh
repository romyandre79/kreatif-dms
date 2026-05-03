#!/usr/bin/env bash
set -e

OUT_DIR="bin"
mkdir -p "$OUT_DIR"

TARGET=${1:-""}

if [ -z "$TARGET" ]; then
    echo "[INFO] No platform specified, building for current host..."
    GOOS=$(go env GOOS)
    GOARCH=$(go env GOARCH)
    TARGET_DIR="$OUT_DIR"
elif [ "$TARGET" == "all" ]; then
    ./build_all.sh
    exit 0
else
    IFS="/" read -r GOOS GOARCH <<< "$TARGET"
    if [ -z "$GOARCH" ]; then
        echo "[ERROR] Invalid platform format. Use GOOS/GOARCH (e.g., linux/amd64)"
        exit 1
    fi
    TARGET_DIR="$OUT_DIR/${GOOS}_${GOARCH}"
    mkdir -p "$TARGET_DIR"
fi

SUFFIX=""
if [ "$GOOS" == "windows" ]; then
    SUFFIX=".exe"
fi

echo "[BUILD] $GOOS/$GOARCH - Server..."
CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "$TARGET_DIR/server$SUFFIX" ./cmd/server
echo "[OK] $TARGET_DIR/server$SUFFIX"

echo "[BUILD] $GOOS/$GOARCH - Worker..."
CGO_ENABLED=0 GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "$TARGET_DIR/worker$SUFFIX" ./cmd/worker
echo "[OK] $TARGET_DIR/worker$SUFFIX"

echo ""
echo "[DONE] Build complete. Binaries in ./$TARGET_DIR/"
