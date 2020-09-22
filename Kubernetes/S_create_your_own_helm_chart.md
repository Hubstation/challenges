# Create your own Helm Chart

## Expectation

Deploy an instance of the dockercloud/hello-world container using a Helm Chart you created

## Cluster Required

In order to complete this challenge, you'll need access to a Kubernetes cluster. 

If you don't have one, you can get access to a free namespace in [Okteto Cloud](https://cloud.okteto).

## Solution

In order to create a helm chart, first thing we need to do is [install the `helm` CLI](https://helm.sh/docs/intro/install/).

You can run the `create` command to create a basic chart, with all the required files:

```
$ helm create my-challenge
```

This will create the following files:

```
$ ls -R my-challenge
```

```
Chart.yaml  charts      templates   values.yaml

./charts:

./templates:
NOTES.txt           deployment.yaml     ingress.yaml        serviceaccount.yaml
_helpers.tpl        hpa.yaml            service.yaml        tests

./templates/tests:
test-connection.yaml
```

To deploy the application, run the `install` command: 

```
$ helm install dev my-challenge
```

```
NAME: dev
LAST DEPLOYED: Wed Sep 30 20:45:32 2020
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
1. Get the application URL by running these commands:
  export POD_NAME=$(kubectl get pods --namespace shopware -l "app.kubernetes.io/name=my-challenge,app.kubernetes.io/instance=dev" -o jsonpath="{.items[0].metadata.name}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl --namespace default port-forward $POD_NAME 8080:80
```

For extra points, try to figure out how to change the image (hint, take a look at the `values.yaml` file).