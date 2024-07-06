FROM golang:1.21 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . ./

WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /core-engine

# Deploy the application binary into a lean image
FROM alpine:3.19 AS build-release-stage

# Install Git
RUN apk add --no-cache git

WORKDIR /

COPY --from=build-stage /core-engine /core-engine

ENTRYPOINT ["/core-engine"]
