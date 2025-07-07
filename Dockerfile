FROM golang:1.24.3

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o authapp

COPY wait-for-postgres.sh .
RUN chmod +x wait-for-postgres.sh

EXPOSE 8080

CMD ["./wait-for-postgres.sh", "./authapp"]