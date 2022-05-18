FROM golang as builder

RUN mkdir /build

ADD . /build

WORKDIR /build
RUN GOOS=linux GOARCH=amd64 go build -o hasura ./cmd/hasura/main.go

FROM alpine

COPY --from=builder /build/hasura /hasura
ENTRYPOINT ["/hasura"]
