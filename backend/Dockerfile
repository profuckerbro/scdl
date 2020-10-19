## We specify the base image we need for our
## go application
FROM golang:latest
## We create an /app directory within our
## image that will hold our application source
## files
RUN go get github.com/profuckerbro/scdl
WORKDIR $GOPATH/src/github.com/profuckerbro/scdl

RUN go install

RUN scdl
