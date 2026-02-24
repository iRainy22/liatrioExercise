FROM ubuntu:24.04


RUN apt-get update && apt-get install -y golang-go ca-certificates


WORKDIR /app


COPY . .


RUN go mod download

RUN go build -o server

EXPOSE 3000

CMD ["./server"]
