## We specify the base image we need for our
## go application
FROM golang:latest
## We create an /app directory within our
## image that will hold our application source
## files

WORKDIR /backend
COPY backend ./


RUN go install

RUN scdl
