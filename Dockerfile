############# builder
FROM golang:1.14.2 AS builder

WORKDIR /go/src/github.com/gardener/gardener-extension-networking-cilium
COPY . .
RUN make install

############# gardener-extension-networking-cilium
FROM alpine:3.11.3 AS gardener-extension-networking-cilium

COPY charts /charts
COPY --from=builder /go/bin/gardener-extension-networking-cilium /gardener-extension-networking-cilium
ENTRYPOINT ["/gardener-extension-networking-cilium"]
