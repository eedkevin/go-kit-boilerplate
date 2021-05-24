FROM golang:1.16-alpine AS builder

ENV GO111MODULE=on
    
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app echo-service/cmd/echo/main.go

FROM alpine:3.7
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 5040 5050 5060
CMD ["./app"]
