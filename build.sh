#MacOS
CGO_ENABLED=1 GOOS=darwin go build -o ./bin/app_darwin

#Linux
CGO_ENABLED=1 GOOS=linux go build -o ./bin/app_linux

#NOTE: since we need CGO_ENABLED=1, if you want to build for windows, you need to install mingw
#https://sourceforge.net/projects/mingw-w64/

#Windows | CXX -> C++ compiler | CC -> C compiler | Without Console Window
CGO_ENABLED=1 GOOS=windows CXX=i686-w64-mingw32-g++ CC=i686-w64-mingw32-gcc GOARCH=386 go build -o ./bin/app.exe -ldflags -H=windowsgui