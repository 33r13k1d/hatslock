set GOOS=windows

set GOARCH=386
go build -ldflags "-s -w -H=windowsgui" -o bin\hatslock.exe

set GOARCH=amd64
go build -ldflags "-s -w -H=windowsgui" -o bin\hatslock64.exe