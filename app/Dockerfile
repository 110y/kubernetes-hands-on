FROM golang:1.13.7-alpine3.11 AS builder

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/110y/kubernetes-hands-on/app

COPY . .
RUN go build -o /usr/bin/app .

# runtime image
FROM alpine:3.11.3

COPY --from=builder /usr/bin/app /usr/bin/app
RUN apk update
RUN apk add --no-cache ca-certificates

EXPOSE 8080

CMD ["/usr/bin/app"]
