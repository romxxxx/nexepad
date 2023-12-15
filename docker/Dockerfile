# -- multistage docker build: stage #1: build stage
FROM golang:1.19-alpine AS build

RUN mkdir -p /go/src/github.com/romxxxx/nexepad

WORKDIR /go/src/github.com/romxxxx/nexepad

RUN apk add --no-cache curl git openssh binutils gcc musl-dev

COPY go.mod .
COPY go.sum .


# Cache nexepad dependencies
RUN go mod download

COPY . .

RUN go build $FLAGS -o nexepad .

# --- multistage docker build: stage #2: runtime image
FROM alpine
WORKDIR /app

RUN apk add --no-cache ca-certificates tini

COPY --from=build /go/src/github.com/romxxxx/nexepad/nexepad /app/
COPY --from=build /go/src/github.com/romxxxx/nexepad/infrastructure/config/sample-nexepad.conf /app/

USER nobody
ENTRYPOINT [ "/sbin/tini", "--" ]
