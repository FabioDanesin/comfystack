#!/bin/bash 
#
# Script interno per inizializzazione dei servizi docker
#

cd /comfystack/server/

# Installa Go
apt update && apt install golang ca-certificates curl -y 
# Download delle dipendenze
go mod download


cd static
# Download dei pacchetti usati per il frontend
curl -L https://cdn.jsdelivr.net/npm/bulma@latest/css/bulma.min.css > bulma.css # Libreria css
curl -L https://unpkg.com/htmx.org@latest/dist/htmx.min.js          > htmx.min.js # Libreria HTMX
cd ..

# Start server
go run main.go