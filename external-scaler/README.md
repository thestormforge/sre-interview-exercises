# external-scaler exercise

We, at StormForge, are fans of autoscaling in kubernetes. HPA already offers a basic set of capabilities which when combined with things like [Keda](https://keda.sh/) allow for many more possibilities.

In this exercise, you will help us use our internal license service which keeps track of our customers subscription to scale up/down the backend services corresponding to their tenant.

For your convenience, a kubernetes manifest directory has already been created. We are using `kustomize` to tie it all together. Let's start with deploying what we have got so far:

1. Install [kind](https://kind.sigs.k8s.io/) on your system, if you don't already have a working kubernetes cluster.
2. With `kubectl` and `kustomize` installed atop, perform the following:
    ```shell
    kustomize build k8s/licenses | k apply -f -

    # namespace names below correspond to tenant IDs in ./k8s/licenses/sample-data.json
    kubectl create namespace tenant-a || kubectl -n tenant-a apply -f ./k8s/tenant-specific/scaled-service.yaml
    kubectl create namespace tenant-b || kubectl -n tenant-b apply -f ./k8s/tenant-specific/scaled-service.yaml
    kubectl create namespace tenant-c || kubectl -n tenant-c apply -f ./k8s/tenant-specific/scaled-service.yaml
    ```
3. Now, if you paid attention to the sample data, `tenant-b` doesn't have a working license. Their `scaled-service` should be scaled down to 0 replicas until an action is taken by the support team to reinstate their license. You need to write a keda external scaler that utilizes the existing license service to acheive this. A few requirements:
    1. The scaler should react dynamically to a change in licenses. For now, our JSON based license service isn't very dynamic (it requires a configmap change and restart), , but still, your scaler should be reactive to changing data and scale tenants up and down as needed.
    2. The scaler should respond appropriately to license service having any hiccups. That is, it should wait on making any scaling decisions if license service is in an error state. (You can simulate such troubles by scaling license service to 0.)
    3. When a license gets reinstated, the number of replicas for a tenant should be back to the default of 2.

## Implementation and Building

1. The gRPC protos provided by Keda have been pre-built using `protoc` for your convenience [here](./pkg/externalscaler) so you can go straight to writing against their contract.
    1. [grpcurl](https://github.com/fullstorydev/grpcurl) might be helpful for validating your scaler but don't let that discourage you from writing tests that run as part of `go test`.
2. Once you are done writing the code for the scalar (there is plenty of help [here](https://keda.sh/docs/2.8/concepts/external-scalers/)), you can either modify the goreleaser config and push your docker image to a public GHCR under your GitHub, or use the local option in kind documented [here](https://kind.sigs.k8s.io/docs/user/local-registry/).


## Deployment

Get the k8s manifests written and running for your scaler from under the [k8s](./k8s) folder before you introduce Keda `ScaledObject` CRDs for the scaled-service.
