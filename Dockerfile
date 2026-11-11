############# builder
FROM golang:1.27.4@sha256:08bcb99a38f85a265f9a93d1d3323e2440eec333380f3ea26ee0dec9264bb41c AS builder

WORKDIR /go/src/github.com/gardener/gardener-extension-networking-cilium

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG EFFECTIVE_VERSION

RUN make install EFFECTIVE_VERSION=$EFFECTIVE_VERSION

############# gardener-extension-networking-cilium
FROM gcr.io/distroless/static-debian13:nonroot@sha256:25ac8f0c3b2334a9eb05452266bd4697864a8c8ef0ef71e5e6480d7f92868329 AS gardener-extension-networking-cilium
WORKDIR /

COPY --from=builder /go/bin/gardener-extension-networking-cilium /gardener-extension-networking-cilium
ENTRYPOINT ["/gardener-extension-networking-cilium"]

############# gardener-extension-admission-cilium
FROM gcr.io/distroless/static-debian13:nonroot@sha256:25ac8f0c3b2334a9eb05452266bd4697864a8c8ef0ef71e5e6480d7f92868329 AS gardener-extension-admission-cilium
WORKDIR /

COPY --from=builder /go/bin/gardener-extension-admission-cilium /gardener-extension-admission-cilium
ENTRYPOINT ["/gardener-extension-admission-cilium"]
