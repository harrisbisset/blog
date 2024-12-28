export PATH := $(PWD)/bin:$(PATH)

.PHONY: pre-render build run-pre-render build-pre-render build-frontend

run: run-pre-render run-frontend
build: build-pre-render build-frontend	

build-pre-render:
	cd backend/pre-render && \
	go mod tidy && \
	go build .

run-pre-render:
	./backend/pre-render/pre-render
	cp -r ./blog_posts/published/ ./frontend/render/posts/
	cp -r ./blog_posts/rendered/ ./frontend/render/posts/

build-frontend:
	cd frontend && templ generate && go mod tidy && \
	npx tailwindcss -c tailwind.config.js -i main.css -o render/dist/tailwind.min.css --minify
	cd frontend && go build -o bin/ .

run-frontend:
	cd frontend && bin/frontend

# build-backend-client:
# 	cd backend/client && templ generate && go mod tidy && \
# 	npx tailwindcss -c tailwind.config.js -i main.css -o render/dist/tailwind.min.css --minify
# 	cd backend && go build -o bin/ .