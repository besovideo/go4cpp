#!/bin/bash

g++ -fPIC -shared -o libgo4c.so ../go4c.cpp

go build -ldflags "-w -s" -o hello.app ../main.go
./hello.app
