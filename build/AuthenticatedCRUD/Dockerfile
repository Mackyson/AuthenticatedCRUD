FROM golang:alpine as BUILD
MAINTAINER "Mackyson <mackyson72@gmail.com>"
COPY . /go/src/app
WORKDIR /go/src/app/
RUN go build ./cmd/AuthenticatedCRUD/main.go

FROM alpine:latest as RUN
COPY --from=BUILD /go/src/app/main .
CMD ["./main"]
