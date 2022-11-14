# Halyard to Spinnaker Operator migration tool

This tool is used to migrate your current Halyard-generated halconfig to the spinnaker operator kustomize files format. This works best with a currently working halconfig directory. If you are having issues running this properly, please ensure that you are able to deploy your halconfig with Halyard prior to running this CLI. This CLI does not validate that your Halconfig will deploy. Only that you have the correct format structure.



### Installation
To install this tool, download it to either your local machine or your Docker container where Halyard is installed.


MacOS
```
$ bash -c 'curl -L https://github.com/austinthao5/golang_proto_test/releases/latest/download/binary.tgz | tar -xz'
```



Linux (If you're running Halyard on a Docker container, use this option)
```
$ bash -c 'curl -L https://github.com/austinthao5/golang_proto_test/releases/latest/download/binary.tgz | tar -xz'
```
<br />
To use the CLI once it's downloaded, run


```
$ cd binary/macOS
```
OR
```
$ cd binary/linux
```
<br />


Then run the CLI like this
```
$ ./armory-config convert --help
```

Once the configuration has been converted to a Spinnaker Kustomize format, install the Operator

OSS:

https://github.com/armory/spinnaker-operator

Armory:

https://docs.armory.io/armory-enterprise/installation/armory-operator


Once Operator has been installed, you can apply your new Kustomize files to Kubernetes using the following command


```
kubectl apply -k ./output_directory -n spinnaker
```

### Usage
This will output the Kustomize files to a local directory. It's a one way conversion from Halyard -> Operator meaning that you cannot provide a set of Spinnaker Kustomize files and convert them to halconfig.

The default namespace that Kustomize expects you to deploy this to is the `spinnaker` namespace. If you want to change this, edit the `Kustomization.yml` file that gets outputted. On `line 4` there is a namespace field that can be changed to a different namespace.

**Required Flags**
- `--halconfig`: Provide the entire Hal directory where your halconfig lives.

**Optional Flags**
- `--output`: Provide the output directory where your Kustomize files will be generated (default: `~./operatorConfig`)
- `--spin_flavor`: Specify whether you want to convert the Halconfig to an OSS or Armory distribution (default: `ARMORY`)
- `--override_deployment`: Override the currentDeployment field in Halyard if you have multiple Spinnaker configurations in the same config file. (default: `default`)
- `--skip_validations`: Whether or not to skip validating the halconfig files. (default: `N`)
- `--writefiles`: Whether or not the CLI will copy all local files referenced in the halconfig and output them into the `files-patch.yml` file. (default: `false`)

**Example usage**
`armory-config convert --halconfig ~/.hal --output ./test_output`

**Validation**

To validate whether or not the outputted Kustomize files will render properly, you can run this command
`kustomize build ./output_directory`

To validate whether or not the Kustomize files will deploy properly to Kubernetes you can run this command
`kubectl apply -k /output_directory -n spinnaker --server-dry-run`

### Generate Proto Files
`protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/proto/deploymentConfigurations/providers/AppEngine.proto`

### Learn more about the Spinnaker Operator
- [Spinnaker Operator Docs](https://docs.armory.io/armory-enterprise/installation/armory-operator)
- [Spinnaker Operator Repo](https://github.com/armory/spinnaker-operator)
- [K8s Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)
- [Armory support help](https://support.armory.io/support)
- [Manual migration for OSS](https://github.com/armory/spinnaker-operator/blob/master/doc/migrate.md)
- [Manual migration for Armory](https://docs.armory.io/armory-enterprise/installation/armory-operator/hal-op-migration/)