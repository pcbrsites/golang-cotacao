FROM golang:1.24-alpine AS build

WORKDIR /app
COPY go.mod ./


RUN apk add --no-cache gcc musl-dev

COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -o server .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/server .

RUN apk add --no-cache sqlite-libs

CMD ["./server"]
