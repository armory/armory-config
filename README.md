# Halyard to Spinnaker Operator migration tool

This tool is used to migrate your current Halyard-generated halconfig to the Spinnaker Operator Kustomize files format. This works best with a currently working halconfig directory. If you are having issues running this properly, please ensure that you are able to deploy your halconfig with Halyard prior to running this CLI. This CLI does not validate that your Halconfig will deploy. Only that you have the correct format structure.

### Limitations
- This CLI does NOT install Operator for you.
- This CLI does NOT deploy anything. It only converts the config files into a Spinnaker Operator Kustomize format.
- `canary.serviceIntegrations` has an issue with AWS accounts when using the `endpoint` field. This field accepts a `string` input whereas other `endpoint` fields for Prometheus, Datadog, etc expect an `object` input. To workaround this, comment or remove this line from your configuration and migrate it manually by copy/pasting the value into your Kustomize Files. The rest of the Canary AWS account configuration can be safely converted.
-  `--writefiles` does not write files from `persistentStorage.gcs.jsonPath`. Only from `providers.kubernetes.kubeconfigFile`, `authz.requiredGroupMembership.google.credentialsPath`, and `pubsub.google.subscriptions.templatePath`.
- Your profiles and service-settings files will NOT be validated. The CLI does it's best to try and paste these files into the correct areas however, if the indentation or format of these files are invalid, this won't cause the CLI to break but, you will likely need to validate the configurations are correct yourself.


### Installation
To install this tool, download it to either your local machine or your Docker container where Halyard is installed.

Create a new directory
```
mkdir armory-config && cd armory-config
```

Download the CLI


MacOs
```
bash -c 'curl -L https://github.com/armory/armory-config/releases/latest/download/armory-config-MacOS.tar.gz | tar -xz'
cd macOS
```



Linux (If you're running Halyard on a Docker container, use this option)
```
bash -c 'curl -L https://github.com/armory/armory-config/releases/latest/download/armory-config-linux.tar.gz | tar -xz'
cd linux
```
<br />

Then run the CLI like this
```
./armory-config convert --help
```
<br />


After the configuration has been converted to a Spinnaker Operator Kustomize format, install the Operator

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

The default namespace that the CLI expects you to deploy these files to is the `spinnaker` namespace. If you want to change this, edit the `Kustomization.yml` file that gets outputted. On `line 4` there is a namespace field that can be changed to a different namespace.

**Required Flags**
- `--halconfig`: Provide the entire Hal directory where your halconfig lives.

**Optional Flags**
- `--output`: Provide the output directory where your Kustomize files will be generated (default: `./operatorConfig`)
- `--spin_flavor`: Specify whether you want to convert the Halconfig to an OSS or Armory distribution (default: `ARMORY`)
- `--override_deployment`: Override the currentDeployment field in Halyard if you have multiple Spinnaker configurations in the same config file. (default: `default`)
- `--skip_validations`: Whether or not to skip validating the halconfig files. (default: `N`)
- `--writefiles`: Whether or not the CLI will copy all local files referenced in the halconfig and output them into the `files-patch.yml` file. (default: `false`)

**Example usage**

```
# Convert the configs using the CLI
armory-config convert --halconfig ~/.hal --output ./test_output --writefiles

# Apply the configuration with Operator
kubectl apply -k ./test_output -n spinnaker
```

**Validation**

To validate whether or not the outputted Kustomize files will render properly, you can run this command

```
kustomize build ./output_directory
```

To validate whether or not the Kustomize files will deploy properly to Kubernetes you can run this command after Operator has been installed

```
kubectl apply -k /output_directory -n spinnaker --server-dry-run
```

### Generate Proto Files


This is the command we used to generate the proto files. Only useful if you want to change any of the proto structs. If you don't know what that means, you can safely ignore this.


```
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/proto/deploymentConfigurations/providers/AppEngine.proto
```

### Learn more about the Spinnaker Operator
- [Spinnaker Operator Docs](https://docs.armory.io/armory-enterprise/installation/armory-operator)
- [Spinnaker Operator Repo](https://github.com/armory/spinnaker-operator)
- [K8s Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)
- [Armory support help](https://support.armory.io/support)
- [Manual migration for OSS](https://github.com/armory/spinnaker-operator/blob/master/doc/migrate.md)
- [Manual migration for Armory](https://docs.armory.io/armory-enterprise/installation/armory-operator/hal-op-migration/)
