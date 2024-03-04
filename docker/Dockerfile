# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/Hoosat-Oy/htnd

WORKDIR /go/src/github.com/Hoosat-Oy/htnd

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache htnd dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o htnd .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/Hoosat-Oy/htnd/htnd /app/
COPY --from=build /go/src/github.com/Hoosat-Oy/htnd/infrastructure/config/sample-htnd.conf /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
