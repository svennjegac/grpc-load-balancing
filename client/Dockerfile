FROM golang:1.12.6-alpine3.10

WORKDIR /app
RUN apk add build-base git bash pkgconfig
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o client -a -tags static_all ./cmd/

CMD ./client
