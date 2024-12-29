FROM golang:1.22

RUN apt-get update && apt-get install -y redis-server \
    && rm -rf /var/lib/apt/lists/*

RUN sed -i 's/^daemonize no/daemonize yes/' /etc/redis/redis.conf

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

EXPOSE 80

COPY . .

RUN go build -o backend .

CMD redis-server /etc/redis/redis.conf && ./backend
