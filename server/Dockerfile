FROM golang:1.23.1

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/cespare/reflex

EXPOSE 8080

CMD reflex-go '*.go' go run main.go --start-service