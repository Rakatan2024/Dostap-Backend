#!/bin/bash

#cd /path/to/your/go/project
pkill -f "go run main.go"
git pull origin main --rebase
go run ./cmd/main.go
#go bin ./cmd/main.go
#pkill bin
#nohup ./bin &
