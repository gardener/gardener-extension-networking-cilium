############# builder
FROM golang:1.23.0 AS builder

WORKDIR /go/src/github.com/gardener/gardener-extension-networking-cilium

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG EFFECTIVE_VERSION

RUN make install EFFECTIVE_VERSION=$EFFECTIVE_VERSION

############# gardener-extension-networking-cilium
FROM gcr.io/distroless/static-debian11:nonroot AS gardener-extension-networking-cilium
WORKDIR /

COPY --from=builder /go/bin/gardener-extension-networking-cilium /gardener-extension-networking-cilium
ENTRYPOINT ["/gardener-extension-networking-cilium"]

############# gardener-extension-admission-cilium
FROM gcr.io/distroless/static-debian11:nonroot AS gardener-extension-admission-cilium
WORKDIR /

COPY --from=builder /go/bin/gardener-extension-admission-cilium /gardener-extension-admission-cilium
ENTRYPOINT ["/gardener-extension-admission-cilium"]
