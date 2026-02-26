#--- Build Stage ---
FROM golang:1.25-alpine AS builder

WORKDIR /app
#copy dependencies and install
COPY api/go.mod api/go.sum ./
RUN go mod download

#copies over the server and builds the binary
COPY api/ ./
RUN go build -o /bin/server ./


#--- Deployment ---
FROM alpine:latest

#copies the binary from the builder
COPY --from=builder /bin/server /bin/server

#opens port 80
EXPOSE 80

#starts the server
CMD ["./bin/server"]
