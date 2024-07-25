run: build
	./bin/main.exe

build:
	go build -ldflags="-s -w" -o bin/main.exe cmd/api/main.go

secret:
	go run cmd/secret/main.go

tailwind:
	tailwindcss -i style.css -o public/css/style.css -m -w
