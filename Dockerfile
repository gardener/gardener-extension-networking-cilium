############# builder
FROM eu.gcr.io/gardener-project/3rd/golang:1.16.2 AS builder

WORKDIR /go/src/github.com/gardener/gardener-extension-networking-cilium
COPY . .
RUN make install

############# gardener-extension-networking-cilium
FROM eu.gcr.io/gardener-project/3rd/alpine:3.13.2 AS gardener-extension-networking-cilium

COPY charts /charts
COPY --from=builder /go/bin/gardener-extension-networking-cilium /gardener-extension-networking-cilium
ENTRYPOINT ["/gardener-extension-networking-cilium"]
