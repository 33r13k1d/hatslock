set GOOS=windows

set GOARCH=386
go build -ldflags -H=windowsgui -o bin\capslang.exe

set GOARCH=amd64
go build -ldflags -H=windowsgui -o bin\capslang64.exe