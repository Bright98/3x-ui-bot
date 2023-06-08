FROM golang:1.20-buster
RUN apt-get update && apt-get install -y git
#RUN apk add git
## We create an /app directory within our
## image that will hold our application source
## files
ENV GOPROXY=https://goproxy.io,direct
ENV GOPRIVATE=git.mycompany.com,github.com/my/private
RUN mkdir /app
WORKDIR /app
#COPY ams_core/go.mod ams_core/go.sum ./
#RUN go mod download
ADD ./3x-ui-bot ./

#RUN go mod init
RUN go list
RUN go mod download

RUN go build -o main .

EXPOSE 9053
## Our start command which kicks off
## our newly created binary executable
CMD "/app/main"