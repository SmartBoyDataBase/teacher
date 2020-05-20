FROM golang:1.12-alpine as builder
RUN apk add git
COPY . /go/src/sbdb-teacher
ENV GO111MODULE on
WORKDIR /go/src/sbdb-teacher
RUN go get && go build

FROM alpine
MAINTAINER longfangsong@icloud.com
COPY --from=builder /go/src/sbdb-teacher/sbdb-teacher /
WORKDIR /
CMD ./sbdb-teacher
ENV PORT 8000
EXPOSE 8000