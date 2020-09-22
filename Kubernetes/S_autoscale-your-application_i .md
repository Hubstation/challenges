# Scale up your application when it receives traffic

## Expectation

The number of replicas of your application automatically increases when the container CPU exceeds 50%

## Cluster Required

In order to complete this challenge, you'll need access to a Kubernetes cluster. 

If you don't have one, you can get access to a free namespace in [Okteto Cloud](https://cloud.okteto).

## Solution

Kubernetes has a resource called [Horizontal Pod Autoscaler](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/). We can use that to tell kubernetes to automatically scale our application based on metrics like CPU and Memory. 

First, let's deploy a hello-world application:

```
kubectl apply -f https://raw.githubusercontent.com/okteto/go-autoscale/master/k8s.yml
```

This is a basic application that takes a request on port 8080 and makes a long mathematical operation.

To enable autoscaling on an application, you can use the `autoscale` command in `kubectl`. In this case, let's configure it so it can scale to a max of 3 pods, whenever the CPU exceeds 50%.

```
$ kubectl autoscale deployment calculate --cpu-percent=50 --min=1 --max=3
```

Now, we will see how the autoscaler reacts to increased load. We will start a container, and send an infinite loop of queries to the calculate service (please run it in a different terminal):

```
$ kubectl run -it --rm load-generator --image=busybox /bin/sh
```

```
Hit enter for command prompt

while true; do wget -q -O- http://calculate:8080; done
```

Wait a minute or so, and check the status of the scaler:

```
$ kubectl get hpa
```

```
NAME          REFERENCE                TARGETS   MINPODS   MAXPODS   REPLICAS   AGE
calculate     Deployment/calculate     286%/50%   1         3        3          
```

You can learn more about autoscaling application [in this tutorial](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/).