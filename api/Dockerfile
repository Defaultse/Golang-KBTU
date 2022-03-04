FROM golang:1.16-alpine

WORKDIR /api

COPY *.go ./

RUN go build -o .

CMD [ "/cmd/main" ]