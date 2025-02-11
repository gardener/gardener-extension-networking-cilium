---
apiVersion: apps/v1
kind: DaemonSet
metadata:
    name: host-network
    namespace: {{ .HelmDeployNamespace }}
    labels:
      app: host-network
spec:
spec:
  selector:
    matchLabels:
      app: host-network
  template:
    metadata:
      labels:
        app: host-network
    spec:
      hostNetwork: true
      containers:
      - name: host-network-test-connectivity
        image: "ubuntu"
        command: 
        - /bin/bash
        - -c
        - |
          apt-get update && apt-get install -y curl iputils-ping; /script/network-test.sh
        volumeMounts:
        - name: networking-test
          mountPath: /script
      volumes:
      - name: networking-test
        configMap:
          defaultMode: 511
          name: network-test
---        
apiVersion: apps/v1
kind: DaemonSet
metadata:
    name: pod-network
    namespace: {{ .HelmDeployNamespace }}
    labels:
      app: pod-network
spec:
  selector:
    matchLabels:
      app: pod-network
  template:
    metadata:
      labels:
        app: pod-network
    spec:
      containers:
      - name: pod-network-test-connectivity
        image: "ubuntu"
        command: 
        - /bin/bash
        - -c
        - |
          apt-get update && apt-get install -y curl iputils-ping; /script/network-test.sh
        volumeMounts:
        - name: networking-test
          mountPath: /script
      volumes:
      - name: networking-test
        configMap:
          defaultMode: 511
          name: network-test
---
kind: Service
apiVersion: v1
metadata:
  name: host-network
  namespace: {{ .HelmDeployNamespace }}
spec:
  selector:
    app: host-network
  ports:
  - name: test-port
    port: 80
---
kind: Service
apiVersion: v1
metadata:
  name: pod-network
  namespace: {{ .HelmDeployNamespace }}
spec:
  selector:
    app: pod-network
  ports:
  - name: test-port
    port: 80
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: network-test-rbac
subjects:
  - kind: ServiceAccount
    name: default
    namespace: default
roleRef:
  kind: ClusterRole
  name: cluster-admin
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: network-test
  namespace: {{ .HelmDeployNamespace }}
data:
  network-test.sh: |
    curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
    chmod +x ./kubectl
    mv ./kubectl /usr/local/bin/kubectl
    
    function ping_ips() {
        local service_name=$1 
        echo "Testing $service_name"
        ips=$(kubectl get endpoints "$service_name" -o jsonpath='{.subsets[*].addresses[*].ip}')
        for ip in $ips; do 
            ip=$(echo "$ip" | tr -d '"')
            ping -c 1 "$ip" > /dev/null
            if [ $? -ne 0 ]
            then
                echo "ERROR: ping $ip failed"
                exit 1
            fi  
            echo "$(date): $ip - succeeded"
        done
    }
    
    while true
    do 
        sleep 15
        ping_ips "host-network"
        ping_ips "pod-network"
    done
