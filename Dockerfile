FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main cmd/main.go

FROM alpine

COPY --from=builder /app/main /main

EXPOSE 50050

CMD ["/main"]