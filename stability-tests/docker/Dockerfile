ARG nexePAD_IMAGE
ARG nexePAMINER_IMAGE

FROM ${nexePAD_IMAGE} as nexepad
FROM ${nexePAMINER_IMAGE} as nexelliaminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/romxxxx/nexepad

WORKDIR /go/src/github.com/romxxxx/nexepad

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=nexepad /app/ /app/
COPY --from=nexelliaminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/romxxxx/nexepad/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]