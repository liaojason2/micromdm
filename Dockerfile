FROM golang:1.17 as builder

WORKDIR /go/src/github.com/micromdm/micromdm/

ARG TARGETARCH
ARG TARGETOS

ENV CGO_ENABLED=0 \
	GOARCH=$TARGETARCH \
	GOOS=$TARGETOS

COPY . .

RUN make deps
RUN make


FROM alpine:latest

RUN apk --update add ca-certificates

COPY --from=builder /go/src/github.com/micromdm/micromdm/build/linux/micromdm /usr/bin/
COPY --from=builder /go/src/github.com/micromdm/micromdm/build/linux/mdmctl /usr/bin/

EXPOSE 80 443 2195 2196 2197 5223

VOLUME ["/var/db/micromdm"]

CMD ["micromdm", "serve"]