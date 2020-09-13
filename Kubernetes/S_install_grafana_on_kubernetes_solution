#Go to KataKoda Playground
https://www.katacoda.com/courses/kubernetes/playground
# Install Arkade
```
master $ curl -sLS https://dl.get-arkade.dev | sudo sh

x86_64
Downloading package https://github.com/alexellis/arkade/releases/download/0.6.12/arkade as /tmp/arkade
Download complete.

Running with sufficient permissions to attempt to move arkade to /usr/local/bin
New version of arkade installed to /usr/local/bin
Creating alias 'ark' for 'arkade'.
            _             _
  __ _ _ __| | ____ _  __| | ___
 / _` | '__| |/ / _` |/ _` |/ _ \
| (_| | |  |   < (_| | (_| |  __/
 \__,_|_|  |_|\_\__,_|\__,_|\___|

Get Kubernetes apps the easy way

Version: 0.6.12
Git Commit: 0415b5fa9d0a6740feb3d9093b7555d38c7e1a51
master $
master $ arkade install grafana
2020/09/13 17:22:56 Client: x86_64, Linux
2020/09/13 17:22:56 User dir established as: /root/.arkade/
https://get.helm.sh/helm-v3.2.4-linux-amd64.tar.gz
/tmp/linux-amd64 linux-amd64/
/tmp/helm linux-amd64/helm
/tmp/README.md linux-amd64/README.md
/tmp/LICENSE linux-amd64/LICENSE
2020/09/13 17:22:57 extracted tarball into /tmp: 3 files, 0 dirs (580.473276ms)
Downloaded to:  /root/.arkade/bin/helm helm
"stable" has been added to your repositories

VALUES values.yaml
Command: /root/.arkade/bin/helm [upgrade --install grafana stable/grafana --namespace grafana --version 5.0.4 --values /tmp/charts/grafana/values.yaml --set sidecar.datasources.enabled=false --set sidecar.dashboards.enabled=false]
Release "grafana" does not exist. Installing it now.
NAME: grafana
LAST DEPLOYED: Sun Sep 13 17:23:03 2020
NAMESPACE: grafana
STATUS: deployed
REVISION: 1
NOTES:
1. Get your 'admin' user password by running:

   kubectl get secret --namespace grafana grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

2. The Grafana server can be accessed via port 80 on the following DNS name from within your cluster:

   grafana.grafana.svc.cluster.local

   Get the Grafana URL to visit by running these commands in the same shell:

     export POD_NAME=$(kubectl get pods --namespace grafana -l "app=grafana,release=grafana" -o jsonpath="{.items[0].metadata.name}")
     kubectl --namespace grafana port-forward $POD_NAME 3000

3. Login with the password from step 1 and the username: admin
#################################################################################
######   WARNING: Persistence is disabled!!! You will lose your data when   #####
######            the Grafana pod is terminated.                            #####
#################################################################################
=======================================================================
=                      grafana has been installed                     =
=======================================================================


# Get the admin password:

  kubectl get secret --namespace grafana grafana -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

# Expose the service via port-forward:

  kubectl --namespace grafana port-forward service/grafana 3000:80

# Enable persistence:

  arkade install grafana --persistence



Thanks for using arkade!
```

# Port-forward and access UI 
```
kubectl --namespace grafana --address 0.0.0.0 port-forward service/grafana 3000:80

```

![](https://raw.githubusercontent.com/Hubstation/challenges/master/Kubernetes/images/grafana.png)
