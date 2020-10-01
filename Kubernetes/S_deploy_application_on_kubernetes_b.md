# Deploy an application in a Kubernetes cluster

## Expectation

Deploy an application that's accessible outside of your cluster.

## Cluster Required

In order to complete this challenge, you'll need access to a Kubernetes cluster. 

If you don't have one, you can get access to a free namespace in [Okteto Cloud](https://cloud.okteto.com).

## Solution 1

Deploy an application that includes a deployment, and a service configured with to use a Load Balancer. Your application will be accessible via the LoadBalancer's IP.

```
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-world
spec:
  replicas: 1
  selector:
    matchLabels:
      app: hello-world
  template:
    metadata:
      labels:
        app: hello-world
    spec:
      containers:
      - image: okteto/hello-world:golang
        name: hello-world
```

```
apiVersion: v1
kind: Service
metadata:
  name: hello-world
spec:
  type: LoadBalancer
  ports:
  - name: "hello-world"
    port: 8080
  selector:
    app: hello-world
```

## Solution 2

Install a cloud native ingress (e.g. NGINX-Ingress or Traeffix) in your cluster, and then create a deployment, a service and an ingress object. Your application will be accessible via the Ingress' IP.

```
# deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hello-world
spec:
  replicas: 1
  selector:
    matchLabels:
      app: hello-world
  template:
    metadata:
      labels:
        app: hello-world
    spec:
      containers:
      - image: okteto/hello-world:golang
        name: hello-world
```

```
apiVersion: v1
kind: Service
metadata:
  name: hello-world
spec:
  type: ClusterIP
  ports:
  - name: "hello-world"
    port: 8080
  selector:
    app: hello-world
```

```
# ingress.yaml

apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: nginx
  name: ingress
spec:
  rules:
  - host: your-host.example.com
    http:
      paths:
      - backend:
          serviceName: web
          servicePort: 80
        path: /
```