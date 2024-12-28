.PHONY: run build build-frontend build-backend run-frontend run-backend

run: run-frontend run-backend
build: build-frontend build-backend

run-backend:
	@cd frontend && bin/backend

run-frontend:
	@cd frontend && bin/frontend

build-frontend:
	cd frontend && templ generate && go mod tidy && \
	npx tailwindcss -c tailwind.config.js -i main.css -o render/dist/tailwind.min.css --minify
	cd frontend && go build -o bin/ .

build-backend:
	cd backend && templ generate && go mod tidy && \
	npx tailwindcss -c tailwind.config.js -i main.css -o render/dist/tailwind.min.css --minify
	cd backend && go build -o bin/ .