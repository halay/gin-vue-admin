@echo off
setlocal
echo please build web 
pause
set GOOS=linux
cd .\server
echo start build server...
go build -o ./bin/deedmx
copy config.yaml bin\
cd ..\web
echo start build web...
REM npm run build
cd ..\
set GOOS=windows
set SERVER=ec2-3-107-8-41.ap-southeast-2.compute.amazonaws.com
set USER=admin
set TEMP_PATH=/tmp/
set REMOTE_PATH=/data/deedmx/
set REMOTE_DIST_PATH=/data/deedmx/dist/
set PEM_FILE=deedmx.pem

echo start push file...
ssh -i "%PEM_FILE%" %USER%@%SERVER% "sudo -i && systemctl stop deedmx"
scp -i "%PEM_FILE%" .\server\bin\deedmx "%USER%@%SERVER%:%TEMP_PATH%"
scp -i "%PEM_FILE%" .\server\bin\config.yaml "%USER%@%SERVER%:%TEMP_PATH%"

ssh -i "%PEM_FILE%" %USER%@%SERVER% "sudo -i && rm -rf %REMOTE_DIST_PATH%/*"
scp -i "%PEM_FILE%"  -r .\web\dist "%USER%@%SERVER%:%TEMP_PATH%"

ssh -i "%PEM_FILE%"  %USER%@%SERVER% "sudo -i && mv %TEMP_PATH%deedmx %REMOTE_PATH% && mv %TEMP_PATH%config.yaml %REMOTE_PATH% && mv %TEMP_PATH%dist %REMOTE_PATH% "

ssh -i "%PEM_FILE%"  %USER%@%SERVER% "sudo -i && systemctl start deedmx"

if %errorlevel% equ 0 (
    echo success
) else (
    echo error
)

endlocal
pause