# external-scaler exercise

At StormForge, we are fans of autoscaling in Kubernetes. The horizontal pod scaler in Kubernetes offers a basic set of capabilities, but when combined with a tool like [KEDA](https://keda.sh/) we can achieve more control and precision when scaling workloads.

The objective of this exercise is to use KEDA to check an API endpoint and based on the returned value, make a decision to scale up or down another backend service. For your convenience, a set of Kubernetes manifests representing the services have already been created. 

To start, follow these steps to deploy the services;

1. Install a Kubernetes cluster. [kind](https://kind.sigs.k8s.io/) or [minikube](https://minikube.sigs.k8s.io/docs/start/) or [Docker Desktop](https://docs.docker.com/desktop/kubernetes/) will definitely make this easy.
2. With `kubectl` and `kustomize` installed atop, perform the following:
    ```shell
    kustomize build k8s/licenses | kubectl apply -f -

    # namespace names below correspond to tenant IDs in ./k8s/licenses/sample-data.json
    kubectl create namespace tenant-a || kubectl -n tenant-a apply -f ./k8s/tenant-specific/scaled-service.yaml
    kubectl create namespace tenant-b || kubectl -n tenant-b apply -f ./k8s/tenant-specific/scaled-service.yaml
    kubectl create namespace tenant-c || kubectl -n tenant-c apply -f ./k8s/tenant-specific/scaled-service.yaml
    ```
3. Now, if you paid attention to the sample data, `tenant-b` doesn't have a working license. Their `scaled-service` should be scaled down to 0 replicas until an action is taken by the support team to reinstate their license. You need to write a KEDA external scaler that utilizes the existing license service to acheive this. A few requirements:
    1. The scaler should react dynamically to a change in licenses. For now, our JSON based license service isn't very dynamic (it requires a configmap change and restart), but still, your scaler should be reactive to changing data and scale tenants up and down as needed.
    2. The scaler should respond appropriately to license service having any hiccups. That is, it should wait on making any scaling decisions if license service is in an error state. (You can simulate such troubles by scaling license service to 0.)
    3. When a license gets reinstated, the number of replicas for a tenant should be back to the default of 2.

## Implementation and Building

1. The gRPC protos provided by KEDA have been pre-built using `protoc` for your convenience [here](./pkg/externalscaler) so you can go straight to writing against their contract.
    1. [grpcurl](https://github.com/fullstorydev/grpcurl) might be helpful for validating your scaler but don't let that discourage you from writing tests that run as part of `go test`.
2. Once you are done writing the code for the scaler (there is plenty of help [here](https://keda.sh/docs/2.8/concepts/external-scalers/)), you can either modify the goreleaser config and push your Docker image to a public GHCR under your GitHub, or use the local option in kind documented [here](https://kind.sigs.k8s.io/docs/user/local-registry/).


## Deployment

Get the Kubernets manifests written and running for your scaler from under the [k8s](./k8s) folder before you introduce KEDA `ScaledObject` CRDs for the scaled-service.
