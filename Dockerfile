FROM golang:latest

# Specify working dir for container
WORKDIR /app

# Add env vars
ENV APP_DB_USERNAME=postgres
ENV APP_DB_NAME=postgres
ENV APP_DB_PASSWORD=mysecretpassword
ENV APP_DB_HOST=host.docker.internal
ENV APP_DB_PORT=5432
ENV APP_DB_SSLMODE=disable 

# Add relevant files to image
COPY . .
RUN go mod download

# Build go app, name it "main"
RUN go build -o main .

# Expose port to receive requests
EXPOSE 8010

# Run app
CMD ["./main"]