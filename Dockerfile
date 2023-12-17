# syntax=docker/dockerfile:1

FROM golang:1.21.5 AS build

WORKDIR /blog

# TODO: clean up
COPY go.mod go.sum ./
COPY ./shared/ ./shared/
COPY ./tem/ ./tem/
RUN go mod download

COPY ./blogPosts/ ./blogPosts/
COPY *.go ./
RUN go get ./shared/
RUN go get ./tem/
RUN go install


RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main


FROM alpine:latest as final

WORKDIR /blog

RUN mkdir /blog/index
COPY --from=build /blog/main ./
COPY --from=build /blog/blogPosts ./blogPosts/

ENTRYPOINT ["/blog/main"]
EXPOSE 80
