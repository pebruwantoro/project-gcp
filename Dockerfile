FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server .

FROM alpine:latest

WORKDIR /root/

# Accept build arguments
ARG PORT
ARG VERSION

RUN echo "PORT=${PORT}" > .env && echo "VERSION=${VERSION}" >> .env

COPY --from=build /app/server .

EXPOSE 8080

CMD ["./server"]