ARG KASPAD_IMAGE
ARG KASPAMINER_IMAGE

FROM ${KASPAD_IMAGE} as hoosatd
FROM ${KASPAMINER_IMAGE} as hoosatminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/Hoosat-Oy/hoosatd

WORKDIR /go/src/github.com/Hoosat-Oy/hoosatd

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=hoosatd /app/ /app/
COPY --from=hoosatminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/Hoosat-Oy/hoosatd/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]