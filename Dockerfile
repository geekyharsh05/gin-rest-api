FROM golang:1.21.1-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

CMD ["go", "run", "main.go"]
