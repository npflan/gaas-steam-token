FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main ./cmd

WORKDIR /dist

RUN cp /build/main .

FROM alpine as runner

COPY --from=builder /build/main ./app/main

EXPOSE 8080
ENV STEAM_WEB_API_KEY=

CMD ["./app/main"]
