FROM golang:alpine AS builder

WORKDIR /apigateway

# Install git for fetching dependencies
RUN apk update && apk add --no-cache git

COPY go.mod .

RUN go mod download

COPY . .

# Build the binary.
RUN go build -o /go/bin/apigateway cmd/apigateway/main.go

## Build lighter image
FROM alpine:latest
LABEL Name=gio-api-gateway-ms Version=1.0.0

# Copy our static executable.
COPY --from=builder /go/bin/apigateway /apigateway
COPY --from=builder /apigateway/config.json /config.json

EXPOSE 8080

# Run the binary.
ENTRYPOINT /apigateway