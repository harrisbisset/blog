# syntax=docker/dockerfile:1

# Fetch
FROM golang:1.23.4 AS fetch-stage
COPY ./frontend/ /frontend/
RUN go mod download

# Generate
FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /frontend
WORKDIR /frontend
RUN ["templ", "generate"]

# Build
FROM golang:1.22.1 AS build-stage
COPY --from=generate-stage /frontend /frontend/
WORKDIR /frontend
RUN CGO_ENABLED=0 GOOS=linux go build -a .

# Static
FROM alpine:latest AS static-stage
WORKDIR /frontend
COPY --from=build-stage /frontend/ ./
RUN ["apk", "add", "curl"]
RUN ["curl", "-sLO", "https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.17/tailwindcss-linux-x64"]
RUN ["chmod", "+x", "tailwindcss-linux-x64"]
RUN ["mv", "tailwindcss-linux-x64", "tailwindcss"]
RUN ["./tailwindcss", "-c", "/frontend/tailwind.config.js", "-i", "/frontend/main.css", "-o", "/frontend/render/dist/tailwind.min.css", "--minify"]

# Deploy
FROM alpine:latest as final
WORKDIR /frontend
COPY --from=static-stage /frontend/ ./
ENTRYPOINT ["/frontend/frontend"]
EXPOSE 80
EXPOSE 443