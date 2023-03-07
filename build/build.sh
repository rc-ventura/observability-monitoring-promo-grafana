#!/bin/sh

go mod download 
go mod verify
go build -v -o /usr/local/bin/app
