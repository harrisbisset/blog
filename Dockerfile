# syntax=docker/dockerfile:1

# Fetch
FROM golang:1.23.4 AS fetch-stage
WORKDIR /frontend
COPY ./frontend/ /frontend/
COPY ./internal/ /internal/
COPY ./blog_posts/ /blog_posts/
COPY ./backend/pre-render/ /backend/pre-render/

# Generate
FROM ghcr.io/a-h/templ:latest AS generate-stage
WORKDIR /frontend
COPY --chown=65532:65532 --from=fetch-stage /frontend /frontend/
RUN ["templ", "generate"]

# Build
FROM golang:1.23.4 AS build-stage
COPY --from=fetch-stage /internal /internal/

COPY --from=generate-stage /frontend /frontend/
WORKDIR /frontend
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a .

COPY --from=fetch-stage /backend/pre-render /backend/pre-render/
WORKDIR /backend/pre-render
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a .


# Static
FROM alpine:latest AS static-stage
WORKDIR /frontend
COPY --from=build-stage /frontend/ /frontend/
COPY --from=build-stage /backend/pre-render/ /backend/pre-render/
COPY --from=fetch-stage /blog_posts/ /blog_posts/

RUN apk add curl
RUN curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.17/tailwindcss-linux-x64
RUN chmod +x tailwindcss-linux-x64
RUN mv tailwindcss-linux-x64 tailwindcss
RUN ./tailwindcss -c /frontend/tailwind.config.js -i /frontend/main.css -o /frontend/render/dist/tailwind.min.css --minify

WORKDIR /
RUN ./backend/pre-render/pre-render
RUN cp -r /blog_posts/published/ /frontend/render/posts/
RUN cp -r /blog_posts/rendered/ /frontend/render/posts/


# Deploy
FROM alpine:latest AS final
WORKDIR /frontend
COPY --from=static-stage /frontend/ ./
ENTRYPOINT ["/frontend/frontend"]
EXPOSE 80
EXPOSE 443