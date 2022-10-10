# Halyard to Spinnaker Operator migration tool

This tool is to migrate you current Halyard-generated halconfig to the spinnaker operator kustomize files format

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
- [Manual migration](https://github.com/armory/spinnaker-operator/blob/master/doc/migrate.md)
