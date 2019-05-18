# example-full [![Build Status](https://cloud.drone.io/api/badges/yellowmegaman/example-full/status.svg)](https://cloud.drone.io/yellowmegaman/example-full)

### What can you find here:
- incomplete golang app which is able to connect to posgres and is listening on port, responding 404 to all requests, except /hello/
- drone.io pipeline for build/pack
- terraform template to launch GKE kubernetes cluster
- yaml files for Stolon (HA postgres) and golang app deployment on kubernetes

### How-to
#### 0) Any commit to this repo is being processed by CI system on cloud.drone.io. Code is being built and every build  results in docker images with tag `latest` and commit hash on docker.hub / quay
Example:
```
docker pull quay.io/yellowmegaman/example-full:a14df448d47ab39a1ed855d4fc5d487cf997f26d
docker pull yellowmegaman/example-full
```

#### 1) configure terraform to work with GCP
```
provider "google" {
  project      = "titanium-messenger-001"
  region       = "europe-west1"
}
```
Then you can export GOOGLE_CREDENTIALS variable with path to auth.json, and GCE_PROJECT with project name.

#### 2) Apply k8s.tf with terraform. (plan first!;)
Now we have 1.12 kubernetes cluster with 4-node nodepool, which can be auto-scaled to 7 nodes. RBAC enabled.

#### 3) Edit k8s/example-full/example-full-Deployment.yaml so it is using proper container tag
#### 4) Authorize to the k8s cluster and apply the yaml's.
```
k apply -R -f k8s
```
Done!

```
$ k -n example-full get all
NAME                                READY   STATUS    RESTARTS   AGE
pod/example-full-5c4b569845-rm5fn   1/1     Running   6          41h

NAME                   TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)    AGE
service/example-full   ClusterIP   None         <none>        8002/TCP   41h

NAME                           DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/example-full   1         1         1            1           41h

NAME                                      DESIRED   CURRENT   READY   AGE
replicaset.apps/example-full-5c4b569845   1         1         1       41h
$ k -n example-full logs pod/example-full-5c4b569845-rm5fn
Successfully connected to db!
```
