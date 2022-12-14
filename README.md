# armory-config

This CLI tool is used to migrate your current Halyard-generated halconfig to the Spinnaker Operator Kustomize files format. This works best with a currently working halconfig directory. If you are having issues running this properly, please ensure that you are able to deploy your halconfig with Halyard prior to running this CLI. This CLI does not validate that your Halconfig will deploy. Only that you have the correct format structure.

## Limitations
- This CLI does NOT install Operator for you.
- This CLI does NOT deploy anything. It only converts the config files into a Spinnaker Operator Kustomize format.
- This CLI does not validate service-settings and profiles files. It does it's best to try and paste the configurations in the right areas however if the indentation or spacing of any of these files are not valid, the CLI won't break but, the formatting will not be deployable with Operator. Please double check the outputted `service-settings.yml` and `profiles-patch.yml` file to ensure a seamless migration.


## Getting Started
**Installation**

- Create a new directory
```
mkdir armory-config && cd armory-config
```

- Download the CLI


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

**Running the CLI**

```
./armory-config convert --help
```
<br />

**Install Operator**

After the configuration has been converted to a Spinnaker Operator Kustomize format, install the Operator

- OSS: https://github.com/armory/spinnaker-operator

- Armory: https://docs.armory.io/armory-enterprise/installation/armory-operator


**Apply the Spinnaker Kustomize Operator configurations**

Once Operator has been installed, you can apply your new Kustomize files to Kubernetes using the following command


```
kubectl apply -k ./output_directory -n spinnaker
```

## Commands

### Convert
```
./armory-config convert
```

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
./armory-config convert --halconfig ~/.hal --output ./test_output --writefiles

# Apply the configuration with Operator
kubectl apply -k ./test_output -n spinnaker
```

## Validation

To validate whether or not the outputted Kustomize files will render properly, you can run this command

```
kustomize build ./output_directory
```

To validate whether or not the Kustomize files will deploy properly to Kubernetes you can run this command after Operator has been installed

```
kubectl apply -k /output_directory -n spinnaker --server-dry-run
```

## Caveats
- `canary.serviceIntegrations` has an issue with AWS accounts when using the `endpoint` field. This field accepts a `string` input whereas other `endpoint` fields for Prometheus, Datadog, etc expect an `object` input. To workaround this, comment or remove this line from your configuration and migrate it manually by copy/pasting the value into your Kustomize Files. The rest of the Canary AWS account configuration can be safely converted.
-  `--writefiles` does not write files from `persistentStorage.gcs.jsonPath`. Only from `providers.kubernetes.kubeconfigFile`, `authz.requiredGroupMembership.google.credentialsPath`, and `pubsub.google.subscriptions.templatePath`.

## Tips

- The default namespace that the CLI expects you to deploy these files to is the `spinnaker` namespace. If you want to change this, edit the `Kustomization.yml` file that gets outputted. On `line 4` there is a namespace field that can be changed to a different namespace.
- If you want a single manifest instead of a Kustomize format, run `kustomize build /output_directory > spinnaker_manifest.yml` 
- If Halyard is installed on a Docker container and you want to remove the outputted files from Docker, you can copy them to your local machine by running `docker cp <container_name>:/path/to/file /path/to/local`
- If you are unable to use curl and download the CLI directly in Docker, you can download it locally and then run `docker cp /path/to/cli <container_name>:/path/in/docker`

## Generate Proto Files


This is the command we used to generate the proto files. Only useful if you want to change any of the proto structs. If you don't know what that means, you can safely ignore this.


```
protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/proto/deploymentConfigurations/providers/AppEngine.proto
```

## Learn more about the Spinnaker Operator
- [Spinnaker Operator Docs](https://docs.armory.io/armory-enterprise/installation/armory-operator)
- [Spinnaker Operator Repo](https://github.com/armory/spinnaker-operator)
- [K8s Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)
- [Armory support help](https://support.armory.io/support)
- [Manual migration for OSS](https://github.com/armory/spinnaker-operator/blob/master/doc/migrate.md)
- [Manual migration for Armory](https://docs.armory.io/armory-enterprise/installation/armory-operator/hal-op-migration/)
