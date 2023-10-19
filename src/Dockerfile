FROM golang:1.19

EXPOSE 7000

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /application

CMD ["/application"]
