
FROM golang:1.19-alpine
LABEL Name="Rahul" Version=1.1
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /docker-gs-ping

EXPOSE 8090

CMD ["/docker-gs-ping"]