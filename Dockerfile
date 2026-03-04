#--- Build Stage ---
FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS builder

WORKDIR /app
#copy dependencies and install
COPY api/go.mod api/go.sum ./
RUN go mod download

#copies over the server and builds the binary
COPY api/ ./

#Declare buildx variables
arg TARGETOS
arg TARGETARCH

# Builds binary for scratch
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build -trimpath -o server . 

#--- Run-time Stage ---
FROM scratch 

#copies the binary from the builder
COPY --from=builder /app/server /server

#opens port 8080 
EXPOSE 8080

#Switch to non-root user
USER 65532:65532

#starts the server
ENTRYPOINT ["/server"]
