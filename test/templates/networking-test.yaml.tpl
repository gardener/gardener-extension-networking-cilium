---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: networking-test
  namespace: {{ .HelmDeployNamespace }}
spec:
  selector:
    matchLabels:
      app: networking-test
  template:
    metadata:
      labels:
        app: networking-test
    spec:
      containers:
      - image: eu.gcr.io/gardener-project/gardener/cilium-cli:1.1.0
        name: networking-shoot-tests-cilium
        command: ["sh", "-c"]
        args:
        - cilium-cli connectivity test
        securityContext:
          privileged: true