FROM golang:1.22-alpine
RUN go install github.com/swaggo/swag/cmd/swag@latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN swag init --parseDependency  --parseInternal --parseDepth 1 -q -d cmd -o ./cmd/docs
RUN go build -o ./vex

EXPOSE 1597
CMD ./vex
