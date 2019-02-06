FROM golang:latest as builder
WORKDIR /go/src/github.com/hdiomede/status-health
RUN go get -u golang.org/x/text/...
RUN go get -u github.com/labstack/echo/...
RUN go get -u github.com/jasonlvhit/gocron ...
RUN go get -u github.com/parnurzeal/gorequest ...
RUN go get gopkg.in/mgo.v2
RUN go get github.com/imdario/mergo
COPY ./ /go/src/github.com/hdiomede/status-health

WORKDIR /go/src/github.com/hdiomede/status-health/cmd/api
RUN CGO_ENABLED=0 GOOS=linux go install

FROM scratch
COPY --from=0 /go/bin/api .
EXPOSE 8080
CMD ["./api"]