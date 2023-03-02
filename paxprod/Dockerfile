FROM golang:1.20.1

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy
CMD ["air", "-c", ".air.toml"]