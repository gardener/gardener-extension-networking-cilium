module github.com/gardener/gardener-extension-networking-cilium

go 1.13

require (
	github.com/ahmetb/gen-crd-api-reference-docs v0.1.5
	github.com/gardener/gardener v1.1.2
	github.com/gardener/gardener-extensions v1.4.0
	github.com/gardener/gardener-resource-manager v0.12.0
	github.com/go-logr/logr v0.1.0
	github.com/gobuffalo/packr/v2 v2.8.0
	github.com/golang/mock v1.4.1
	github.com/onsi/ginkgo v1.12.0
	github.com/onsi/gomega v1.9.0
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v0.0.6
	github.com/spf13/pflag v1.0.5
	k8s.io/api v0.17.3
	k8s.io/apimachinery v0.17.3
	k8s.io/apiserver v0.17.3 // indirect
	k8s.io/client-go v11.0.1-0.20190409021438-1a26190bd76a+incompatible
	k8s.io/code-generator v0.17.3
	k8s.io/component-base v0.17.3
	sigs.k8s.io/controller-runtime v0.5.0
)
