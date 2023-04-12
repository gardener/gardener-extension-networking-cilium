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
      - image: eu.gcr.io/gardener-project/gardener/cilium-cli:1.2.0
        name: networking-shoot-tests-cilium
        command: ["sh", "-c"]
        args:
        - cilium-cli connectivity test --test '!to-entities-world,!to-fqdns,!client-egress-l7,!client-egress-l7-named-port'
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
