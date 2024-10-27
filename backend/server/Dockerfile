FROM ubuntu:latest

# Cartella applicazione
WORKDIR /comfystack/server/ 
# Installa Go
RUN apt update && apt get golang
# Download delle dipendenze
RUN go mod download
# Copia sorgente ed esegui
COPY . .
ENTRYPOINT [ "go run main.go" ]