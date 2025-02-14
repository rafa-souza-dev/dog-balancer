FROM golang:1.22.5

WORKDIR /app

COPY go.mod ./
RUN go mod download 

COPY . .
RUN go build -v -o dog_balancer

CMD ["./dog_balancer"]
