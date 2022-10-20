package migrate

import (
	"strings"
)

type Providers struct {
	AppEngine      string
	Aws            string
	Ecs            string
	Azure          string
	Dcos           string
	DockerRegistry string
	Google         string
	Huaweicloud    string
	Kubernetes     string
	Tencentcloud   string
	Oracle         string
	Cloudfoundry   string
	Enable         string //Which provider is currently enable
}

func GetProvidersData(KustomizeData Kustomize) string {

	// KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Aws.GetAwsAcc()
	// KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.AppEngine
	prob := Providers{}
	prob.SetProvidersData(KustomizeData)

	str := `
  validation:
    providers:
      ` + prob.Enable + `:
        enabled: true    # Default: true. Indicate if operator should do connectivity checks to configured kubernetes accounts before applying the manifest
  spinnakerConfig:
    config:
      providers:
        appengine:
          ` + prob.AppEngine + `
        aws:
          ` + prob.Aws + `
        ecs:
          ` + prob.Ecs + `
        azure:
          ` + prob.Azure + `
        dcos:
          ` + prob.Dcos + `
        dockerRegistry:
          ` + prob.DockerRegistry + `
        google:
          ` + prob.Google + `
        huaweicloud:
          ` + prob.Huaweicloud + `
        kubernetes:
          ` + prob.Kubernetes + `
        tencentcloud:
          ` + prob.Tencentcloud + `
        oracle:
          ` + prob.Oracle + `
        cloudfoundry:
          ` + prob.Cloudfoundry

	return str
}

// This function fills the Providers struct in a valid format
func (ProvidersData *Providers) SetProvidersData(KustomizeData Kustomize) error {

	// Slice of function calls to fill all the Providers Data
	functionCalls := []func(KustomizeData Kustomize) error{
		ProvidersData.SetAppEngineData,
		ProvidersData.SetAwsData,
		ProvidersData.SetEcs,
		ProvidersData.SetAzure,
		ProvidersData.SetDcos,
		ProvidersData.SetDockerRegistry,
		ProvidersData.SetGoogle,
		ProvidersData.SetHuaweicloud,
		ProvidersData.SetKubernetes,
		ProvidersData.SetTencentcloud,
		ProvidersData.SetOracle,
		ProvidersData.SetCloudfoundry,
	}

	for _, function := range functionCalls {
		err := function(KustomizeData)
		if err != nil {
			return err
			// return fmt.Errorf("Error while executing the function:%s Error:(%s)", function, err)
		}
	}

	return nil
}

func (ProvidersData *Providers) SetAppEngineData(KustomizeData Kustomize) error {

	str := `enabled: false
	accounts: []`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.AppEngine.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.AppEngine = str

	return nil
}

func (ProvidersData *Providers) SetAwsData(KustomizeData Kustomize) error {

	str := `enabled: true
	primaryAccount: ` + /* KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Aws.PrimaryAccount +*/ `                # Must be one of the configured AWS accounts
	# accounts: []
	  accounts:
	  - name: XXXXX
	  accountId: "XXXXX"            # (Required). Your AWS account ID to manage. See the AWS IAM User Guide for more information.
	  assumeRole: XXXXX           # (Required). If set, will configure a credentials provider that uses AWS Security Token Service to assume the specified role. Example: “user/spinnaker” or “role/spinnakerManaged”
	  lifecycleHooks: []                   # (Optional). Configuration for AWS Auto Scaling Lifecycle Hooks. For more information, see: https://docs.aws.amazon.com/autoscaling/ec2/userguide/lifecycle-hooks.html
	  permissions: {}
	  providerVersion: V1
	  regions:
	  - name: us-east-1
	  - name: us-east-2
	  - name: us-west-1
	  - name: us-west-2
	  bakeryDefaults:                        # Configuration for Spinnaker’s image bakery.Configuration for Spinnaker’s image bakery.
	    baseImages: []
	  accessKeyId: ` + /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Aws.AccessKey +*/ `      # Only needed if cluster worker nodes don't have IAM roles for talking to the target aws account
	  secretAccessKey: ` + /*KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Aws.SecretAccessKey +*/ `  # Only needed if cluster worker nodes don't have IAM roles for talking to the target aws account
	  defaultKeyPairTemplate: '` + /*KustomizeData.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Halyard.Providers.Aws.DefaultKeyPairTemplate +*/ `'
	  defaultRegions:
	  - name: us-east-1
	  - name: us-east-2
	  - name: us-west-1
	  - name: us-west-2
	  defaults:
	    iamRole: BaseIAMRole
	  features:
	    cloudFormation:
	      enabled: true                       # (Default: false). Enable cloudformation support on this AWS account.`

	str = strings.Replace(str, "\t", "          ", -1)

	// log.Println("str:" + str)

	ProvidersData.Aws = str

	return nil
}

func (ProvidersData *Providers) SetEcs(KustomizeData Kustomize) error {

	str := `enabled: true
	accounts:
	- name: halyard-migrator
	  requiredGroupMembership: []
	  providerVersion: V1
	  permissions: {}
	  awsAccount: XXXXX                   # Must be one of the configured AWS accounts
	primaryAccount: halyard-migrator             # Must be one of the configured AWS ECS accounts`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Ecs.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Ecs = str

	return nil
}

func (ProvidersData *Providers) SetAzure(KustomizeData Kustomize) error {

	str := `enabled: false
	accounts: []
	bakeryDefaults:
	  templateFile: azure-linux.json
	  baseImages: []`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Azure.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Azure = str

	return nil
}

func (ProvidersData *Providers) SetDcos(KustomizeData Kustomize) error {

	str := `enabled: false
	accounts: []
	clusters: []`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Dcos.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Dcos = str

	return nil
}

func (ProvidersData *Providers) SetDockerRegistry(KustomizeData Kustomize) error {

	str := `enabled: true
	accounts:
	- name: halyard-migrator
	  requiredGroupMembership: []
	  permissions: {}
	  address: XXXXX
	  username: XXXXX
	  password: XXXXX
	  email: fake.email@spinnaker.io
	  cacheIntervalSeconds: 3600
	  clientTimeoutMillis: 60000
	  cacheThreads: 1
	  paginateSize: 100
	  sortTagsByDate: false
	  trackDigests: false
	  insecureRegistry: false
	  repositories: []
	primaryAccount: halyard-migrator`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.DockerRegistry.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.DockerRegistry = str

	return nil
}

func (ProvidersData *Providers) SetGoogle(KustomizeData Kustomize) error {

	str := `enabled: false
	accounts: []
	bakeryDefaults:
	  templateFile: gce.json
	  baseImages: []
	  zone: us-central1-f
	  network: default
	  useInternalIp: false`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Google.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Google = str

	return nil
}

func (ProvidersData *Providers) SetHuaweicloud(KustomizeData Kustomize) error {

	str := `enabled: false
	accounts: []
	bakeryDefaults:
	  baseImages: []`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Huaweicloud.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Huaweicloud = str

	return nil
}

func (ProvidersData *Providers) SetKubernetes(KustomizeData Kustomize) error {

	//TODO check which one is enable
	if true {
		ProvidersData.Enable = "kubernetes"
	}

	str := `enabled: true
	accounts:
	- name: halyard-migrator
	  requiredGroupMembership: []
	  permissions: {}
	  dockerRegistries: []
	  providerVersion: V2
	  context: XXXXX
	  configureImagePullSecrets: true
	  cacheThreads: 1
	  namespaces: []
	  omitNamespaces: []
	  kinds: []
	  omitKinds: []
	  customResources: []
	  cachingPolicies: []
	  kubeconfigFile: /home/spinnaker/.kube/config
	  rawResourcesEndpointConfig:
		kindExpressions: []
		omitKindExpressions: []
	  oAuthScopes: []
	  onlySpinnakerManaged: false
	primaryAccount: halyard-migrator`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Kubernetes.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Kubernetes = str

	return nil
}

func (ProvidersData *Providers) SetTencentcloud(KustomizeData Kustomize) error {

	str := `enabled: false
	accounts: []
	bakeryDefaults:
	  baseImages: []`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Tencentcloud.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Tencentcloud = str

	return nil
}

func (ProvidersData *Providers) SetOracle(KustomizeData Kustomize) error {

	str := `enabled: false
	accounts: []
	bakeryDefaults:
	  templateFile: oci.json
	  baseImages: []`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Oracle.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Oracle = str

	return nil
}

func (ProvidersData *Providers) SetCloudfoundry(KustomizeData Kustomize) error {

	str := `enabled: false
	accounts: []`

	// enabled: ` + KustomizeData.Halyard.DeploymentConfiguration[KustomizeData.CurrentDeploymentPos].Providers.Cloudfoundry.Enable + `

	str = strings.Replace(str, "\t", "          ", -1)
	ProvidersData.Cloudfoundry = str

	return nil
}
