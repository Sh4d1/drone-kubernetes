# Drone Kubernetes 
[![Build Status](https://cloud.drone.io/api/badges/Sh4d1/drone-kubernetes/status.svg)](https://cloud.drone.io/Sh4d1/drone-kubernetes) [![](https://images.microbadger.com/badges/image/sh4d1/drone-kubernetes.svg)](https://hub.docker.com/r/sh4d1/drone-kubernetes/ "Get your own image badge on microbadger.com")

Drone plugin to create/update Kubernetes resources.

It uses the latest k8s go api, so it is intened to use on Kubernetes 1.9+. I can't guarantee it will work for previous versions.

You can directly pull the image from [sh4d1/drone-kubernetes](https://hub.docker.com/r/sh4d1/drone-kubernetes/)
## Supported resources
Currently, this plugin supports:
* apps/v1
  * DaemonSet
  * Deployment
  * ReplicaSet
  * StatefulSet
* apps/v1beta1
  * Deployment
  * StatefulSet
* apps/v1beta2
  * DaemonSet
  * Deployment
  * ReplicaSet
  * StatefulSet
* v1
  * ConfigMap 
  * PersistentVolume 
  * PersistentVolumeClaim 
  * Pod 
  * ReplicationController 
  * Service 
* extensions/v1beta1
  * DaemonSet
  * Deployment
  * Ingress
  * ReplicaSet

## Inspiration 

It is inspired by [vallard](https://github.com/vallard) and his plugin [drone-kube](https://github.com/vallard/drone-kube).


## Usage

Here is how you can use this plugin:
```
pipeline:
  deploy:
    image: sh4d1/drone-kubernetes
    kubernetes_template: deployment.yml
    kubernetes_namespace: default
    secrets: [kubernetes_server, kubernetes_cert, kubernetes_token]
```

## Secrets

You need to define these secrets before.
```
$ drone secret add --image=sh4d1/drone-kubernetes -repository <your-repo> -name KUBERNETES_SERVER -value <your API server>
```
```
$ drone secret add --image=sh4d1/drone-kubernetes -repository <your repo> -name KUBERNETES_CERT -value <your base64 encoded cert>
```
```
$ drone secret add --image=sh4d1/drone-kubernetes -repository <your repo> -name KUBERNETES_TOKEN -value <your token>
```

### How to get values of `KUBERNETES_CERT` and `KUBERNETES_TOKEN`

List secrets of `default` namespace

```
$ kubectl get -n <namespace of secret> default secret
```

Show the `ca.crt` and `token` from secret

```
$ kubectl get secret -n <namespace of secret> <name of your drone secret> -o yaml | egrep 'ca.crt:|token:'
```

You can copy/paste the encoded certificate to the `KUBERNETES_CERT` value.
For the `KUBERNETES_TOKEN`, you need to decode it:

* `echo "<encoded token>" | base64 -d`
* `kubectl describe secret -n <your namespace> <drone secret name> | grep 'token:'`


TODO
