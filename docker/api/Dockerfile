FROM golang:1.23

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src ./src
RUN go build src/main.go

# RUN CGO_ENABLED=0 GOOS=linux go build -o app

EXPOSE 8080

CMD ["/app/main"]