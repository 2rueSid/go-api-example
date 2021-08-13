FROM golang:1.16.7-alpine3.14 as builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .

RUN go run github.com/prisma/prisma-client-go db push

RUN CGO_ENABLED=0 GOOS=linux go build -o server -a -v .

EXPOSE 5000

ENTRYPOINT ["./server"]