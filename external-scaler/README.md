# external-scaler exercise

We, at StormForge, are fans of autoscaling in kubernetes. HPA already offers a basic set of capabilities which when combined with things like [Keda](https://keda.sh/) allow for many more possibilities.

In this exercise, you will help us use our internal license service which keeps track of our customers subscription to scale up/down the backend services corresponding to their tenant.

For your convenience, a kubernetes manifest directory has already been created. We are using `kustomize` to tie it all together. Let's start with deploying what we have got so far:

1. Install [kind](https://kind.sigs.k8s.io/) on your system, if you don't already have a working kubernetes cluster.
2. With `kubectl` and `kustomize` installed atop, perform the following:
    ```shell
    kustomize build k8s/licenses | k apply -f -

    kubectl create namespace tenant-a || kubectl -n tenant-a apply -f ./k8s/tenant-specific/scaled-service.yaml
    kubectl create namespace tenant-b || kubectl -n tenant-b apply -f ./k8s/tenant-specific/scaled-service.yaml
    kubectl create namespace tenant-c || kubectl -n tenant-c apply -f ./k8s/tenant-specific/scaled-service.yaml
    ```
3. Now, if you paid attention to the sample data, `tenant-b` doesn't have a working license. Their `scaled-service` should be scaled down to 0 replicas until an action is taken by the support team to reinstate their license. You need to write some kind of automation that utilizes the existing license service to acheive this. If you have some ideas, go ahead! If you find yourself in a need of hint or two, let us know.
