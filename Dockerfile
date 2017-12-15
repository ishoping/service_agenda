FROM golang:1.8

RUN mkdir -p /go/src/service_agenda
RUN mkdir -p /go/src/service_agenda/data
RUN mkdir -p /go/bin
Run mkdir -p /go/pkg
RUN apt-get update && apt-get install -y --no-install-recommends \
		apt-utils \
		sqlite3

RUN go get -d -v \
	github.com/codegangsta/negroni \
	github.com/go-xorm/xorm \
	github.com/gorilla/mux \
	github.com/mattn/go-sqlite3 \
	github.com/spf13/pflag \
	github.com/unrolled/render


VOLUME ["/go/src/service_agenda/data"]
ADD . /go/src/service_agenda

WORKDIR /go/src/service_agenda

RUN go install -o server .

WORKDIR cli
RUN go install -o cli .

WORKDIR /go/bin

EXPOSE 8080

CMD ["/go/src/service_agenda/server"]