# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY . /app
RUN go mod download

RUN go build -o /hr_api

EXPOSE 5000

CMD [ "/hr_api" ]