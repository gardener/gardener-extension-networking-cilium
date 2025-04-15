---
apiVersion: batch/v1
kind: Job
metadata:
  name: network-test
  namespace: {{ .HelmDeployNamespace }}
  labels:
    app: networking-test
spec:
  template:
    metadata:
      labels:
        app: networking-test
    spec:
      containers:
      - image: europe-docker.pkg.dev/gardener-project/releases/gardener/cilium-cli:1.10.0
        name: networking-shoot-tests-cilium-port-forward
        command: [ "sh", "-c" ]
        args:
          - cilium-cli hubble port-forward
        securityContext:
          capabilities:
            add:
              - NET_ADMIN
        env:
          - name: KUBECONFIG
            value: /etc/kubeconfig/kubeconfig
        volumeMounts:
          - name: shoot-kubeconfig
            mountPath: "/etc/kubeconfig"
            readOnly: true
      - image: europe-docker.pkg.dev/gardener-project/releases/gardener/cilium-cli:1.10.0
        name: networking-shoot-tests-cilium
        command: ["sh", "-c"]
        args:
        - cilium-cli connectivity test --test '!to-entities-world,!to-fqdns,!client-egress-l7,!client-egress-l7-named-port,!client-egress-tls-sni,!check-log-errors'
        securityContext:
          capabilities:
            add:
              - NET_ADMIN
        env:
        - name: KUBECONFIG
          value: /etc/kubeconfig/kubeconfig
        volumeMounts:
        - name: shoot-kubeconfig
          mountPath: "/etc/kubeconfig"
          readOnly: true
      volumes:
      - name: shoot-kubeconfig
        secret:
          secretName: kubeconfig
      restartPolicy: Never
  backoffLimit: 0
