FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o server .

FROM alpine:latest

WORKDIR /root/

COPY .env.example ./.env

COPY --from=build /app/server .

EXPOSE 8080

CMD ["./server"]