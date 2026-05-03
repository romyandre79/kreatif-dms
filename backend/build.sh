#!/usr/bin/env bash
set -euo pipefail

OUT_DIR="bin"
mkdir -p "$OUT_DIR"

echo "Building server..."
go build -o "$OUT_DIR/server" ./cmd/server
echo "[OK] $OUT_DIR/server"

echo "Building worker..."
go build -o "$OUT_DIR/worker" ./cmd/worker
echo "[OK] $OUT_DIR/worker"

echo ""
echo "Build complete. Binaries in ./$OUT_DIR/"
