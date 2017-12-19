FROM golang
MAINTAINER Karthik Gaekwad

RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/prometheus/client_golang/prometheus/promhttp
COPY . /go/src/github.com/karthequian/prometheus-demo
WORKDIR /go/src/github.com/karthequian/prometheus-demo
RUN go get && go build -o /bin/hello

ENTRYPOINT ["/bin/hello"]