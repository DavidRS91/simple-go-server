FROM golang:latest

WORKDIR /app

ENV APP_DB_USERNAME=postgres
ENV APP_DB_NAME=postgres
ENV APP_DB_PASSWORD=mysecretpassword
ENV APP_DB_HOST=host.docker.internal
ENV APP_DB_PORT=5432
ENV APP_DB_SSLMODE=disable 

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 8010

CMD ["./main"]