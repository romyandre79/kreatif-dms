#!/usr/bin/env bash
set -euo pipefail

VENV_DIR="venv"
ENTRY="main.py"
APP_NAME="ocr-service"

echo "============================================"
echo " OCR Service Build"
echo "============================================"

# --- Setup virtual environment ---
if [ ! -f "$VENV_DIR/bin/activate" ]; then
    echo "Creating virtual environment..."
    python3 -m venv "$VENV_DIR"
    echo "[OK] venv created"
else
    echo "[OK] venv already exists"
fi

source "$VENV_DIR/bin/activate"

# --- Install dependencies ---
echo "Installing dependencies..."
pip install --upgrade pip -q
pip install -r requirements.txt
echo "[OK] Dependencies installed"

# --- Optional: bundle with PyInstaller ---
if [ "${1:-}" = "--bundle" ]; then
    echo "Bundling with PyInstaller..."
    pip install pyinstaller -q
    pyinstaller --onefile --name "$APP_NAME" \
        --hidden-import=uvicorn.logging \
        --hidden-import=uvicorn.loops \
        --hidden-import=uvicorn.loops.auto \
        --hidden-import=uvicorn.protocols \
        --hidden-import=uvicorn.protocols.http \
        --hidden-import=uvicorn.protocols.http.auto \
        --hidden-import=uvicorn.lifespan \
        --hidden-import=uvicorn.lifespan.on \
        "$ENTRY"
    echo "[OK] Executable: dist/$APP_NAME"
else
    echo ""
    echo "Build complete. Run with:"
    echo "  source $VENV_DIR/bin/activate"
    echo "  python $ENTRY"
    echo ""
    echo "To create a standalone executable:"
    echo "  ./build.sh --bundle"
fi
