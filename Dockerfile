############# builder
FROM golang:1.18.3 AS builder

WORKDIR /go/src/github.com/gardener/gardener-extension-networking-cilium
COPY . .
RUN make install

############# gardener-extension-networking-cilium
FROM gcr.io/distroless/static-debian11:nonroot AS gardener-extension-networking-cilium
WORKDIR /

COPY charts /charts
COPY --from=builder /go/bin/gardener-extension-networking-cilium /gardener-extension-networking-cilium
ENTRYPOINT ["/gardener-extension-networking-cilium"]
