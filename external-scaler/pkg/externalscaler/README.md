# externalscaler proto generation

If you want to repeat the steps for a newer version in the future:

Download the latest proto [here](https://raw.githubusercontent.com/kedacore/keda/main/pkg/scalers/externalscaler/externalscaler.proto) into your downloads.

```shell
# From the base of this repo, run the following:
brew install protoc-gen-go-grpc

protoc -I ~/Downloads ~/Downloads/externalscaler.proto --go_out=pkg/externalscaler
protoc -I ~/Downloads ~/Downloads/externalscaler.proto --go-grpc_out=pkg/externalscaler
```
