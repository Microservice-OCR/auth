FROM golang:1.21.5-alpine AS builder
RUN apk update && apk add --no-cache git

WORKDIR /go/auth
COPY ./auth .

RUN go install
RUN go build -o /go/auth/bin/auth

FROM scratch
COPY --from=builder /go/auth/bin/auth .
CMD ["/auth"]