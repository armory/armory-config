# Halyard to Spinnaker Operator migration tool

This tool is used to migrate your current Halyard-generated halconfig to the spinnaker operator kustomize files format. This works best with a currently working halconfig file. If you are having issues running this properly, please ensure that you are able to deploy your halconfig with Halyard prior to running this CLI. This CLI does not validate that your Halconfig will deploy. Only that you have the correct format structure.

### Installation
To install this tool, download it to either your local machine or your Docker container where Halyard is installed.

````
$ bash -c 'curl -L https://github.com/austinthao5/golang_proto_test/releases/latest/download/binary.tgz | tar -xz'
```

### Usage

- `go run main.go convert --halconfig </path/to/halconfig> --output </path/to/output/directory>`
- `go run main.go --help`
- `go run main.go convert --help`

### Example usage
`go run main.go convert --halconfig ~/.hal --output ./test_output`

### Generate Proto Files
`protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/proto/deploymentConfigurations/providers/AppEngine.proto`

### Learn more about the Spinnaker Operator
- [Spinnaker Operator Docs](https://docs.armory.io/armory-enterprise/installation/armory-operator)
- [Spinnaker Operator Repo](https://github.com/armory/spinnaker-operator)
- [K8s Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)
- [Armory support help](https://support.armory.io/support)
- [Manual migration for OSS](https://github.com/armory/spinnaker-operator/blob/master/doc/migrate.md)
- [Manual migration for Armory](https://docs.armory.io/armory-enterprise/installation/armory-operator/hal-op-migration/)
