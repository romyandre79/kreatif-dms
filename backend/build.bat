@echo off
setlocal

set MODULE=github.com/kreatif/dms-backend
set OUT_DIR=bin

if not exist %OUT_DIR% mkdir %OUT_DIR%

echo Building server...
go build -o %OUT_DIR%\server.exe .\cmd\server
if %errorlevel% neq 0 (
    echo [ERROR] Server build failed
    exit /b 1
)
echo [OK] server.exe

echo Building worker...
go build -o %OUT_DIR%\worker.exe .\cmd\worker
if %errorlevel% neq 0 (
    echo [ERROR] Worker build failed
    exit /b 1
)
echo [OK] worker.exe

echo.
echo Build complete. Binaries in .\%OUT_DIR%\
endlocal
