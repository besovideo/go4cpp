

g++ -fPIC -shared -o go4c.dll ../go4c.cpp

go build -ldflags "-w -s" -o hello.exe ../main.go
hello.exe

