FROM golang:latest AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o myapp main.go

FROM builder 

WORKDIR /root/

COPY --from=builder /app/myapp .

EXPOSE 3000

CMD ["./myapp"]
