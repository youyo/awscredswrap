FROM golang:1-alpine AS build-env
ENV DIR /go/src/github.com/youyo/awscredswrap
WORKDIR ${DIR}
ADD . ${DIR}
RUN apk add --update --no-cache ca-certificates git
RUN go build -o dist/awscredswrap awscredswrap/main.go

FROM alpine:3
LABEL maintainer "youyo <1003ni2@gmail.com>"
COPY --from=build-env /go/src/github.com/youyo/awscredswrap/dist/awscredswrap /awscredswrap
RUN apk add --update --no-cache ca-certificates bash
ENTRYPOINT ["/awscredswrap"]
