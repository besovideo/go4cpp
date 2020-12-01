#!/bin/bash

g++ -fPIC -shared -o libgo4c.so ../go4c.cpp

export LD_LIBRARY_PATH=.

go build -ldflags "-w -s" -o hello.app ../main.go
./hello.app
