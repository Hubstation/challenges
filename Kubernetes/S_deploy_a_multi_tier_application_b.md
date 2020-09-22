# Deploy a multi-tier application in a Kubernetes cluster

## Expectation

Deploy an application that has at least two components to your Kubernetes cluster.

## Cluster Required

In order to complete this challenge, you'll need access to a Kubernetes cluster. 

If you don't have one, you can get access to a free namespace in [Okteto Cloud](https://cloud.okteto).

## Solution

The Guestbok Application is a great example of a multi-tier application. It includes a redis cache, and a web frontend. 

To launch the redis instance:

```
kubectl apply -f https://github.com/okteto/go-guestbook/blob/master/manifests/redis.yaml
```

To launch the guestbook web frontend:

```
kubectl apply -f https://github.com/okteto/go-guestbook/blob/master/manifests/guestbook.yaml
```

To access the application, start a port-forward on port 8080:

```
kubectl port-forward svc/guestbook 8080:8080
```

And go to http://localhost:8080!