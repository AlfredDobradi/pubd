FROM golang:buster AS build

WORKDIR /build

COPY . /build

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o target/pubd ./cmd/pubd/...

FROM busybox:latest

COPY --from=build /build/target/pubd /usr/bin/pubd

CMD /usr/bin/pubd