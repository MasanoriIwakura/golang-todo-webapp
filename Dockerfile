FROM golang:1.20.3

WORKDIR /app

COPY ./app/go.mod ./
COPY ./app/go.sum ./
RUN go mod download

COPY ./app ./
