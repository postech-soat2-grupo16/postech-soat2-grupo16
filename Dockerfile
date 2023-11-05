FROM golang:1.20

WORKDIR /app
COPY . .
RUN go get -d -v ./...
RUN go build -o build .
EXPOSE 8000

CMD ["./build"]