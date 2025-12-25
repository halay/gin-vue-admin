@echo off
setlocal
echo please build web 
pause
set GOOS=linux
cd .\server
echo start build server...
go build -o ./bin/deedmx
cd ..\
set GOOS=windows
set SERVER=110.41.59.11
set USER=root
set REMOTE_PATH=/data/deedmx/

echo start push file...
ssh %USER%@%SERVER% "systemctl stop deedmx"
scp .\server\bin\deedmx "%USER%@%SERVER%:%REMOTE_PATH%"
ssh %USER%@%SERVER% "systemctl start deedmx"

if %errorlevel% equ 0 (
    echo success
) else (
    echo error
)

endlocal
pause