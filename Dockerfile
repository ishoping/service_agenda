FROM golang:1.8

RUN mkdir -p /go/src/service_agenda
RUN mkdir -p /go/src/service_agenda/data
RUN apt-get update && apt-get install -y --no-install-recommends \
		apt-utils \
		sqlite3

VOLUME ["/go/src/service_agenda/data"]
ADD . /go/src/service_agenda

WORKDIR /go/src/service_agenda

RUN go build -o server .

WORKDIR cli
RUN go build -o cli .

WORKDIR /go/src/service_agenda

EXPOSE 8080

CMD ["/go/src/service_agenda/server"]