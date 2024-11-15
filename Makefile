build:
	templ generate && \
	go mod tidy && \
	npx tailwindcss -c ./tailwind.config.js -i ./render/dist/main.css -o ./render/dist/tailwind.css && \
	go build -o ./bin ./cmd/main.go 