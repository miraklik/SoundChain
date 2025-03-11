FROM golang:1.24-alpine

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/. .
RUN go build -v -o backend-app ./cmd

EXPOSE 8080

CMD ["./backend-app"]