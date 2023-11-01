# Compilation stage
FROM golang:1.20 AS build

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go build -o build .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copies necessary files from the build stage
COPY --from=build /app/build ./build

EXPOSE 8000

CMD ["./build"]
