all: build run

run:
	./bin/main.exe

build:
	go build -ldflags="-s -w" -o bin/main.exe cmd/gospotify/main.go

secret:
	go run cmd/secret/main.go

tailwind:
	tailwindcss -i public/style.css -o public/static/css/style.css --config public/tailwind.config.js --minify --watch
