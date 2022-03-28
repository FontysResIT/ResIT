FROM golang:1.18.0-alpine as builder

ENV GO111MODULE=on

WORKDIR /app
COPY . .

RUN apk --no-cache add git alpine-sdk build-base gcc

RUN go get

RUN go build -o build ./main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/build .
COPY --from=builder /app/config/*.yaml ./config/
CMD ["./build"]
