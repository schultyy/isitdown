FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

CMD [ "/bin/bash" ]

RUN go build

RUN mkdir templates
COPY ./templates/*.html templates/

EXPOSE 8080

CMD [ "./isitdown" ]