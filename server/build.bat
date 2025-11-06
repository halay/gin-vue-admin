@echo off
set GOOS=linux
go build -o ./bin/deedmx
copy config.yaml bin\
set GOOS=windows