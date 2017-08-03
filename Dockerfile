FROM golang:1.8

#copy sources in from host
RUN mkdir -p /go/src/app
WORKDIR /go/src/app
COPY . /go/src/app

RUN go-wrapper download
RUN go-wrapper install
EXPOSE 1337
CMD ["go-wrapper", "run"]
