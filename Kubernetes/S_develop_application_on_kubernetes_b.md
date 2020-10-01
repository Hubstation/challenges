# Deploy and Develop an application directly in a Kubernetes cluster

## Expectation 
Deploy a hello-world application in Kubernetes, launch a development environment with [okteto](https://github.com/okteto/okteto), customize the message and see the result without rebuilding or redeploying.

## Cluster Required

In order to complete this challenge, you'll need access to a Kubernetes cluster. 

If you don't have one, you can get access to a free namespace in [Okteto Cloud](https://cloud.okteto.com).

## Solution

Deploy an existing Kubernetes application:

```
kubectl apply -f https://raw.githubusercontent.com/okteto/go-getting-started/master/k8s.yml
```

Install the latest version of [okteto](https://github.com/okteto/okteto):

### MacOS / Linux
```
 curl https://get.okteto.com -sSfL | sh
```

### Windows

```
Download https://downloads.okteto.com/cli/okteto.exe and add it to your `$PATH`.
```

Create a development environment in your cluster:

```
okteto init
```

Activate your development environment:

```
okteto up
```

Make a change in your local file system, and star the go process on the remote development environment!

```
cindy:hello-world app> go run main.go
```

```
Starting hello-world server...
```

> If you need a hint, [check out this tutorial](https://okteto.com/blog/how-to-develop-go-apps-in-kubernetes/).