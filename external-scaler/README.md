# external-scaler exercise

At StormForge, we are fans of autoscaling in Kubernetes. The horizontal pod scaler in Kubernetes offers a basic set of capabilities, but when combined with a tool like [KEDA](https://keda.sh/) we can achieve more control and precision when scaling workloads.

The objective of this exercise is to use KEDA to check an API endpoint and based on the returned value, make a decision to scale up or down another backend service. For your convenience, a set of Kubernetes manifests representing the services have already been created.

1. Make sure you have a kubernetes cluster. [kind](https://kind.sigs.k8s.io/) or [minikube](https://minikube.sigs.k8s.io/docs/start/) or [Docker Desktop](https://docs.docker.com/desktop/kubernetes/) will definitely make this easy.
2. With `kubectl` installed, perform the following:
    ```shell
    # deploy the license service
    kubectl apply -k k8s/licenses

    # create tenant namespaces (names correspond to tenantIDs in k8s/licenses/sample-data.json)
    foreach ns (tenant-a tenant-b); kubectl create ns $ns; end

    # deploy services in tenant specific namespaces
    kubectl -n tenant-a apply -f k8s/tenant-specific/scaled-service.yaml
    kubectl -n tenant-b apply -f k8s/tenant-specific/scaled-service.yaml
    ```
3. If you look at the sample data in `k8s/licenses/sample-data.json`, `tenant-b` has a license that is expired. Their `scaled-service` should be scaled down to 0 replicas until an action is taken by the support team to reinstate their license.

    This is where you come in, you need to write a KEDA external scaler that checks the license API and will scale the tenant service up or down.

    __Requirements;__

    The scaler should react to any change in the license services. If the license is valid, we scale to 2. If it is invalid, we scale to 0.

## Implementation and Building

1. KEDA uses [gRPC](https://grpc.io/docs/what-is-grpc/introduction/), but the protos provided by KEDA have already been built using `protoc` for your convenience [here](./pkg/externalscaler) so you can go straight to writing against their contract.
    * [grpcurl](https://github.com/fullstorydev/grpcurl) might be helpful for validating your scaler.
2. KEDA doesn't have a plugin for what we need to do, so we need to write an [external scaler](https://keda.sh/docs/2.8/concepts/external-scalers/). If you use the affirmentioned link, most of the work has been done for you.
3. Once you have an external scaler ready, you will need to build the image. You can either modify the goreleaser config and push your Docker image to a public GHCR under your GitHub, or use the local option in kind documented [here](https://kind.sigs.k8s.io/docs/user/local-registry/).
4. The final step is to deploy to your cluster. The keda docs have an example of the manifest for their ScaledObject CRD.

## Testing

Since the sample-data is static, if you make changes, you will have to restart the license service. Taking the license service down can also be a good test in how your scaler reacts to outages.
