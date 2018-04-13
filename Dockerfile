FROM golang:1.10 as builder

#RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/github.com/Sh4d1/drone-kubernetes
WORKDIR /go/src/github.com/Sh4d1/drone-kubernetes

COPY vendor /go/src/github.com/Sh4d1/drone-kubernetes/vendor
COPY Gopkg.toml Gopkg.lock ./
#RUN dep ensure

COPY *.go ./
RUN go build -ldflags "-linkmode external -extldflags -static" -a 

FROM scratch
COPY --from=builder /go/src/github.com/Sh4d1/drone-kubernetes/drone-kubernetes /drone-kubernetes

CMD ["/drone-kubernetes"]
