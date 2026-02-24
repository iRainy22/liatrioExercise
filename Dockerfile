FROM ubuntu:24.04

#Update and install packages 
RUN apt-get update && \
    apt-get install -y \
    golang-go \
    ca-certificates


WORKDIR /app

#copy and install dependencies
COPY api/go.mod api/go.sum .
RUN go mod download

#copy the rest of app, and build it.
COPY api/ .  
RUN go build -o server . 

EXPOSE 3000

CMD ["./server"]
