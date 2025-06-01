APP_NAME ?= app

.PHONY: build
build:
	go build -o ./bin/$(APP_NAME) ./cmd/main.go

.PHONY: css
css:
	tailwindcss -i static/css/input.css -o static/css/output.css

.PHONY: templ
templ:
	templ generate

.PHONY: watch-css
watch-css:
	tailwindcss -i static/css/input.css -o static/css/output.css --watch
