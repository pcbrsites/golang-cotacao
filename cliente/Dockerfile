FROM golang:1.24-alpine AS build

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o client .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/client .


CMD ["./client"]
