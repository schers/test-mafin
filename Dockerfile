FROM golang:1.18.2-alpine3.16 AS builder

RUN apk --update --no-cache add ca-certificates openssl git tzdata gcc musl-dev

ARG cert_location=/usr/local/share/ca-certificates

# Get certificate from "github.com"
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
# Get certificate from "proxy.golang.org"
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# Get certificate from "sum.golang.org"
RUN openssl s_client -showcerts -connect sum.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/sum.golang.crt

RUN update-ca-certificates

WORKDIR /go/src/github.com/schers/test-mafin

ENV GO111MODULE=on

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -a -o build/testmafin ./cmd

FROM alpine:3.16

RUN apk --no-cache add ca-certificates file

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

