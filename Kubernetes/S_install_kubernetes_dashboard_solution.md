#Go to KataKoda Playground
https://www.katacoda.com/courses/kubernetes/playground

# Install Kubernetes Dashboard
```
master $ launch.sh
Waiting for Kubernetes to start...
Kubernetes started
master $ kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.0.0/aio/deploy/recommended.yaml
namespace/kubernetes-dashboard created
serviceaccount/kubernetes-dashboard created
service/kubernetes-dashboard created
secret/kubernetes-dashboard-certs created
secret/kubernetes-dashboard-csrf created
secret/kubernetes-dashboard-key-holder created
configmap/kubernetes-dashboard-settings created
role.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrole.rbac.authorization.k8s.io/kubernetes-dashboard created
rolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created
clusterrolebinding.rbac.authorization.k8s.io/kubernetes-dashboard created
deployment.apps/kubernetes-dashboard created
service/dashboard-metrics-scraper created
deployment.apps/dashboard-metrics-scraper created


```





# Create the admin-user and RBAC role-binding for the user

First, create the user.

You an pipe directly to kubectl.

```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kubernetes-dashboard
EOF

```



Console output

```
master $ cat <<EOF | kubectl apply -f -
> apiVersion: v1
> kind: ServiceAccount
> metadata:
>   name: admin-user
>   namespace: kubernetes-dashboard
> EOF
serviceaccount/admin-user created
```



Then, the role binding for the user. 

```
cat <<EOF | kubectl apply -f -
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kubernetes-dashboard
EOF
```



Console Output:

```
master $ cat <<EOF | kubectl apply -f -
> apiVersion: rbac.authorization.k8s.io/v1
> kind: ClusterRoleBinding
> metadata:
>   name: admin-user
> roleRef:
>   apiGroup: rbac.authorization.k8s.io
>   kind: ClusterRole
>   name: cluster-admin
> subjects:
> - kind: ServiceAccount
>   name: admin-user
>   namespace: kubernetes-dashboard
> EOF
clusterrolebinding.rbac.authorization.k8s.io/admin-user created
```



Then, pull the bearer token from the k8s api.

```
master $ kubectl -n kubernetes-dashboard describe secret $(kubectl -n kubernetes-dashboard get secret | grep admin-user | awk '{print $1}')
Name:         admin-user-token-fqj85
Namespace:    kubernetes-dashboard
Labels:       <none>
Annotations:  kubernetes.io/service-account.name: admin-user
              kubernetes.io/service-account.uid: fa886838-08c7-4332-9cd7-9e095a4cdf4c

Type:  kubernetes.io/service-account-token

Data
====
ca.crt:     1025 bytes
namespace:  20 bytes
token:      eyJhbGciOiJSUzI1NiIsImtpZCI6IlZkWUlLWXZ0Y2dUQmRfQXNJeC1OcnM1b0xoMGw1NzFQc0tYZ2xsWWtIVnMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlcm5ldGVzLWRhc2hib2FyZCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhZG1pbi11c2VyLXRva2VuLWZxajg1Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZXJ2aWNlLWFjY291bnQubmFtZSI6ImFkbWluLXVzZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJmYTg4NjgzOC0wOGM3LTQzMzItOWNkNy05ZTA5NWE0Y2RmNGMiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6a3ViZXJuZXRlcy1kYXNoYm9hcmQ6YWRtaW4tdXNlciJ9.V26rpG0f_h0xWOADnoifDZAmo2Ox7JSAR24vV3WLvMshrkzIXru8I4Sggsp8CNI4zb0Dyn-yqsYAhjyNVcfapqFhIvwCdvtkPFMNrX_I61SFasOl1Jld0xQeVlIZ37BtnZ80uE119JkwsX7Op0syj7gKv_tEh0QX-pWddgdGGP4QBlZ-BQd1xkcF038Se7A_nxlfqdgDntrAo7EFb5WXliw6f5g1sQJDVnIGLP2S-lgRep1LjXcm59nCL3xtZ5cOxU3hLDSSMeD8KwHOBjloGnrca5CgJjQJyZl71OVZvTZ21Jo5oRi5Jt0K6cTPY9n5i8S2fg-UDn_kUP0I9B28Lw
```



Copy the token, then run kubectl port-forward to allow access to the dashboard.



```
master $ kubectl --namespace kubernetes-dashboard port-forward --address 0.0.0.0 service/kubernetes-dashboard 443:443
```



Then you can access the dashboard at port 443.



You will paste the bearer-token into the login prompt.



![kube-dash-login](../../challenges/images/kube-dash-login.png)



And you are presented with the Kuberbetes Dashboard.



![kube-dash-solution](../../challenges/images/kube-dash-solution.png)