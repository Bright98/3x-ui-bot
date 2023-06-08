FROM golang:1.20-buster

#proxy
ENV GOPROXY=https://goproxy.io,direct
ENV GOPRIVATE=git.mycompany.com,github.com/my/private

RUN mkdir /app
WORKDIR /app
#COPY ams_core/go.mod ams_core/go.sum ./

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Build the Go app
RUN go build -o main .

# Run the binary program produced by `go install`
CMD "/app/main"