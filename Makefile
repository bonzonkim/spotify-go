APP_NAME ?= app
build:
	go build -o ./bin/$(APP_NAME) ./cmd/main.go

css:
	tailwindcss -i static/css/input.css -o static/css/output.css

templ:
	templ generate
