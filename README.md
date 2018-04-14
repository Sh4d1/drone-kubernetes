# Drone Kubernetes [![Docker Build Status](https://img.shields.io/docker/build/jrottenberg/ffmpeg.svg)](https://hub.docker.com/r/sh4d1/drone-kubernetes/) [![MicroBadger Size](https://img.shields.io/microbadger/image-size/jumanjiman/puppet.svg)](https://hub.docker.com/r/sh4d1/drone-kubernetes/tags)

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
        secrets: [kubernetes_server, kubernetes_cert, kubernetes_token]
```

## Secrets

You need to define these secrets before.
```
$ drone secret add --image=sh4d1/drone-kubernetes -repository <your-repo> -name KUBERNETES_SERVER -value <your API server>
```
```
$ drone secret add --image=sh4d1/drone-kubernetes -repository <your repo> -name KUBERNETES_CERT <your base64 encoded cert>
```
```
$ drone secret add --image=sh4d1/drone-kubernetes -repository <your repo> -name KUBERNETES_TOKEN <your token>
```

### How to get values of `KUBERNETES_CERT` and `KUBERNETES_TOKEN`

```
$ kubectl get secret -n <namespace of secret> <name of your drone secret> -o yaml | egrep 'ca.crt:|token:'
```

You can copy/paste the encoded certificate to the `KUBERNETES_CERT` value.
For the `KUBERNETES_TOKEN`, you need to decode it:
* `echo "<encoded token> | base64 --decode"`
* `kubectl describe secret -n <your namespace> <drone secret name> | grep 'token:'`



TODO

