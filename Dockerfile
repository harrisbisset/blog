# syntax=docker/dockerfile:1

FROM golang:1.21.5 AS build

WORKDIR /blog


COPY go.mod go.sum tailwind.config.js ./
COPY ./cmd/ ./
COPY ./dist/ ./dist/
COPY ./handler/ ./handler/
COPY ./model/ ./model/
COPY ./posts/MD/ ./posts/MD/
COPY ./public/ ./public/
COPY ./view/ ./view/

RUN ls -la ./model/*
RUN ls -la ./view/*

RUN go mod download
RUN go get ./handler
RUN go get ./view/*
RUN go get ./model/*
RUN go install

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main


FROM alpine:latest as final

WORKDIR /blog

RUN mkdir /blog/index
COPY --from=build /blog/main ./
COPY --from=build /blog/posts/MD/ ./posts/MD/
COPY --from=build /blog/dist/ ./dist/

ENTRYPOINT ["/blog/main"]
EXPOSE 80
