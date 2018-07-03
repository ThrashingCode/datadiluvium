FROM adron/golang-glide

ENV GOPATH /go

ADD . /go/src/github.com/Adron/datadiluvium/
WORKDIR /go/src/github.com/Adron/datadiluvium

RUN cd /go/src/github.com/Adron/datadiluvium && \
    glide install && \
    glide up
