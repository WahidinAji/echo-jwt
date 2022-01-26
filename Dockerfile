FROM golang:1.17.6-buster


WORKDIR /usr/app


COPY . .

RUN go mod download

RUN go build main.go

CMD [ "./main" ]