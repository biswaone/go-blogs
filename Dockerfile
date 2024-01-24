FROM golang:latest

WORKDIR /app

COPY makefile .
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy rest of application code
COPY . .

RUN make build

EXPOSE 8080

CMD ["./bin/goblogs"]
