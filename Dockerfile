FROM golang:1.17

WORKDIR /go/src/app

COPY . .

RUN go get
RUN go install