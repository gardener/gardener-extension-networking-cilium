############# builder
FROM golang:1.27.4@sha256:5e8afd493a7124a8586e134d9b532731161d4b9becf9dee5a5956b107426f93d AS builder

WORKDIR /go/src/github.com/gardener/gardener-extension-networking-cilium

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG EFFECTIVE_VERSION

RUN make install EFFECTIVE_VERSION=$EFFECTIVE_VERSION

############# gardener-extension-networking-cilium
FROM gcr.io/distroless/static-debian13:nonroot@sha256:aafe993665bf6200db38a02c11a9351e782a84b6a45e7f33b6695638c88e3931 AS gardener-extension-networking-cilium
WORKDIR /

COPY --from=builder /go/bin/gardener-extension-networking-cilium /gardener-extension-networking-cilium
ENTRYPOINT ["/gardener-extension-networking-cilium"]

############# gardener-extension-admission-cilium
FROM gcr.io/distroless/static-debian13:nonroot@sha256:aafe993665bf6200db38a02c11a9351e782a84b6a45e7f33b6695638c88e3931 AS gardener-extension-admission-cilium
WORKDIR /

COPY --from=builder /go/bin/gardener-extension-admission-cilium /gardener-extension-admission-cilium
ENTRYPOINT ["/gardener-extension-admission-cilium"]
