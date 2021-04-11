FROM golang:1.16 AS builder

WORKDIR /go/src/k8s-vault-webhook/

COPY go.mod /go/src/k8s-vault-webhook/
COPY go.sum go.sum
RUN go mod download

COPY . /go/src/k8s-vault-webhook/
RUN go get -v -t -d ./... \
    && go build -o k8s-vault-webhook

FROM alpine:latest
ENV DEBUG false
COPY --from=builder /go/src/k8s-vault-webhook/k8s-vault-webhook /usr/local/bin/
RUN apk add --no-cache libc6-compat
ENTRYPOINT ["/usr/local/bin/k8s-vault-webhook"]
