@echo off
setlocal enabledelayedexpansion

set MODULE=github.com/kreatif/dms-backend
set OUT_DIR=bin
if not exist %OUT_DIR% mkdir %OUT_DIR%

:: Check for arguments
set TARGET=%1

if "%TARGET%"=="" (
    echo [INFO] No platform specified, building for current host...
    set GOOS=windows
    set GOARCH=amd64
    set TARGET_DIR=%OUT_DIR%
    set SUFFIX=.exe
) else if "%TARGET%"=="all" (
    call build_all.bat
    exit /b 0
) else (
    for /f "tokens=1,2 delims=/" %%A in ("%TARGET%") do (
        set GOOS=%%A
        set GOARCH=%%B
    )
    if "!GOARCH!"=="" (
        echo [ERROR] Invalid platform format. Use GOOS/GOARCH (e.g., linux/amd64)
        exit /b 1
    )
    
    set SUFFIX=
    if "!GOOS!"=="windows" set SUFFIX=.exe
    
    set TARGET_DIR=%OUT_DIR%\!GOOS!_!GOARCH!
    if not exist !TARGET_DIR! mkdir !TARGET_DIR!
)

echo [BUILD] !GOOS!/!GOARCH! - Server...
set CGO_ENABLED=0
set GOOS=!GOOS!
set GOARCH=!GOARCH!
go build -ldflags="-s -w" -o !TARGET_DIR!\server!SUFFIX! .\cmd\server
if !errorlevel! neq 0 (
    echo [ERROR] Server build failed
    exit /b 1
)
echo [OK] !TARGET_DIR!\server!SUFFIX!

echo [BUILD] !GOOS!/!GOARCH! - Worker...
go build -ldflags="-s -w" -o !TARGET_DIR!\worker!SUFFIX! .\cmd\worker
if !errorlevel! neq 0 (
    echo [ERROR] Worker build failed
    exit /b 1
)
echo [OK] !TARGET_DIR!\worker!SUFFIX!

echo.
echo [DONE] Build complete. Binaries in .\%TARGET_DIR%\
endlocal
