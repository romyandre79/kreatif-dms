@echo off
setlocal

set VENV_DIR=venv
set ENTRY=main.py
set APP_NAME=ocr-service

echo ============================================
echo  OCR Service Build
echo ============================================

REM --- Setup virtual environment ---
if not exist %VENV_DIR%\Scripts\activate.bat (
    echo Creating virtual environment...
    python -m venv %VENV_DIR%
    if %errorlevel% neq 0 (
        echo [ERROR] Failed to create venv. Make sure Python is installed.
        exit /b 1
    )
    echo [OK] venv created
) else (
    echo [OK] venv already exists
)

REM --- Activate and install dependencies ---
call %VENV_DIR%\Scripts\activate.bat

echo Installing dependencies...
pip install --upgrade pip -q
pip install -r requirements.txt
if %errorlevel% neq 0 (
    echo [ERROR] Failed to install dependencies
    exit /b 1
)
echo [OK] Dependencies installed

REM --- Optional: bundle to single executable with PyInstaller ---
if "%1"=="--bundle" (
    echo Bundling with PyInstaller...
    pip install pyinstaller -q
    pyinstaller --onefile --name %APP_NAME% ^
        --hidden-import=uvicorn.logging ^
        --hidden-import=uvicorn.loops ^
        --hidden-import=uvicorn.loops.auto ^
        --hidden-import=uvicorn.protocols ^
        --hidden-import=uvicorn.protocols.http ^
        --hidden-import=uvicorn.protocols.http.auto ^
        --hidden-import=uvicorn.lifespan ^
        --hidden-import=uvicorn.lifespan.on ^
        %ENTRY%
    if %errorlevel% neq 0 (
        echo [ERROR] PyInstaller bundle failed
        exit /b 1
    )
    echo [OK] Executable: dist\%APP_NAME%.exe
) else (
    echo.
    echo Build complete. Run with:
    echo   %VENV_DIR%\Scripts\activate.bat
    echo   python %ENTRY%
    echo.
    echo To create a standalone executable:
    echo   build.bat --bundle
)

endlocal
