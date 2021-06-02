FROM golang:1.16-alpine AS builder

ARG GOPROXY https://proxy.golang.org,direct \
    APP_NAME

ENV GO111MODULE=on \
    GOPROXY=${GOPROXY}
    
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/${APP_NAME}/main.go

FROM alpine:3.7
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 5040 5050 5060
CMD ["./app"]
