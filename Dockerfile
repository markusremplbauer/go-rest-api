FROM golang:1.13-alpine
WORKDIR /app
COPY api/go.mod api/go.sum ./
RUN go mod download
COPY api .
RUN go build -o api .
CMD ["./api"]