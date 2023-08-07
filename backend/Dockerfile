FROM golang:1.20 AS build

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY ./ ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /password-manager ./

FROM build AS run-test
RUN go test -v -race ./...

FROM build AS run-app

EXPOSE 8000

# Run
CMD [ "/password-manager" ]
