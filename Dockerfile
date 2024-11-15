# syntax=docker/dockerfile:1

FROM golang:1.22.6 AS build

WORKDIR /blog

COPY go.mod go.sum tailwind.config.js ./
COPY ./cmd/ ./
COPY ./render/ ./render/
COPY ./components/ ./components/
COPY ./models/ ./models/
COPY ./blog/blog_posts/published/ ./posts/blog_posts/published/

RUN go mod download
RUN go install

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main

FROM alpine:latest as final

WORKDIR /blog

COPY --from=build /blog/main ./
COPY --from=build /blog/blog_posts/published/ ./posts/blog_posts/published/
COPY --from=build /blog/dist/ ./dist/
COPY --from=build /blog/public/ ./public/

ENTRYPOINT ["/blog/main"]
EXPOSE 80