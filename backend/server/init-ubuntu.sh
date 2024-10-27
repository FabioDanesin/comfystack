#!/bin/bash 
cd /comfystack/server/
# Installa Go
apt update && apt install golang ca-certificates -y 
# Download delle dipendenze
go mod download
# Start server
go run main.go