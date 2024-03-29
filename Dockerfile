FROM golang:1.21.5 as builder

ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

# RUN go build -o bin/go-blogs
RUN make build

FROM scratch
COPY --from=builder ["/app/*.env", "/"]
COPY --from=builder ["/app/bin/go-blogs", "/"]

EXPOSE 8080

ENTRYPOINT [ "/go-blogs" ] 