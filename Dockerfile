FROM golang:1.14-alpine3.11 AS builder

RUN apk add --no-cache git gcc musl-dev

WORKDIR /go/src/github.com/schers/test-mafin

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -a -o build/testmafin ./cmd

FROM alpine:3.11

RUN apk add --no-cache ca-certificates file

EXPOSE 8080

ENV HOST=":8080"
ENV STORAGE="/images"

ENV POSTGRES_USER=test
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=test
ENV DB_URL="postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@postgres:5432/$POSTGRES_DB?sslmode=disable"

RUN mkdir -m 755 $STORAGE

COPY --from=builder /go/src/github.com/schers/test-mafin/build/testmafin /usr/local/bin/testmafin
RUN ln -s /usr/local/bin/testmafin /testmafin

ENTRYPOINT ["testmafin"]

