FROM golang:1.18.0-alpine as builder

ENV GO111MODULE=on
WORKDIR /app
COPY . .

RUN apk --no-cache add git alpine-sdk build-base gcc

RUN go get

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init  --parseDependency --parseInternal

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./build

FROM scratch AS final

COPY --from=builder /app/build build
COPY --from=builder /app/docs docs

CMD ["/build"]
