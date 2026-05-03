#!/usr/bin/env bash
set -e

OUT_DIR="bin"
mkdir -p "$OUT_DIR"

PLATFORMS=("windows/amd64" "windows/arm64" "linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64")

echo "[INFO] Starting multi-OS build for Server and Worker..."

for PLATFORM in "${PLATFORMS[@]}"; do
    IFS="/" read -r GOOS GOARCH <<< "$PLATFORM"
    
    SUFFIX=""
    if [ "$GOOS" == "windows" ]; then
        SUFFIX=".exe"
    fi
    
    TARGET_DIR="$OUT_DIR/${GOOS}_${GOARCH}"
    mkdir -p "$TARGET_DIR"
    
    echo "[BUILD] $GOOS/$GOARCH - Server..."
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "$TARGET_DIR/server$SUFFIX" ./cmd/server
    echo "[OK] $TARGET_DIR/server$SUFFIX"
    
    echo "[BUILD] $GOOS/$GOARCH - Worker..."
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o "$TARGET_DIR/worker$SUFFIX" ./cmd/worker
    echo "[OK] $TARGET_DIR/worker$SUFFIX"
done

echo ""
echo "[DONE] Multi-OS build complete. Binaries are in ./$OUT_DIR/"
