package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetAwsData(providersRef *providers.Providers) error {

	// if nil != providersRef.Aws {
	// 	return fmt.Errorf("Aws value is null")
	// }

	//To avoid null pointers
	featuresCloudF := ""
	featuresLambda := ""
	if nil != providersRef.Aws.Features {
		if nil != providersRef.Aws.Features.Cloudformation {
			featuresCloudF = strconv.FormatBool(providersRef.Aws.Features.Cloudformation.Enabled)
		}
		if nil != providersRef.Aws.Features.Lambda {
			featuresLambda = strconv.FormatBool(providersRef.Aws.Features.Lambda.Enabled)
		}
	}

	if providersRef.Aws.Enabled {
		ProvidersData.Enable = "aws"
	}

	str := `enabled: ` + strconv.FormatBool(providersRef.Aws.Enabled) + `
	primaryAccount: ` + providersRef.Aws.PrimaryAccount + `                # Must be one of the configured AWS accounts` +
		GetAwsAccounts(providersRef) +
		GetAwsBakeryDefault(providersRef.Aws.BakeryDefault) + `
	accessKeyId: ` + providersRef.Aws.AccessKey + `      # Only needed if cluster worker nodes don't have IAM roles for talking to the target aws account
	secretAccessKey: ` + providersRef.Aws.SecretAccessKey + `  # Only needed if cluster worker nodes don't have IAM roles for talking to the target aws account
	defaultKeyPairTemplate: '` + providersRef.Aws.DefaultKeyPairTemplate + `'` +
		strings.Replace(getAwsRegions(providersRef.Aws.DefaultRegions, "defaultRegions"), "\t", "    ", -1) + `
	defaults:
	  iamRole: ` + providersRef.Aws.Defaults.IamRole + `
	features:
	  cloudFormation:
	    enabled: ` + featuresCloudF + `                       # (Default: false). Enable cloudformation support on this AWS account.` + `
	  lambda:
	    enabled: ` + featuresLambda

	str = strings.Replace(str, "\t", "          ", -1)

	ProvidersData.Aws = str

	return nil
}

func GetAwsAccounts(provider *providers.Providers) string {
	str := ""

	if nil != provider.Aws.Accounts {
		str += `
		  accounts:`
		for _, account := range provider.Aws.Accounts {
			str += `
		  - name: ` + account.Name + `
		    environment: ` + account.Environment +
				strings.Replace(getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- "), "\t", "   ", -1) +
				getPermissions(account.Permissions) + `
		    defaultKeyPair: ` + account.DefaultKeyPair + `
		    edda : ` + account.Edda + `
		    discovery: ` + account.Discovery + `
		    accountId: '` + account.AccountId + `'` +
				strings.Replace(getAwsRegions(account.Regions, "regions"), "\t", "     ", -1) +
				strings.Replace(getAwsLifecycleHooks(account.LifecycleHooks), "\t", "     ", -1)
			//TODO assumeRole Missing proto
			//TODO providerVersion Missing proto
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func getAwsRegions(regions []*providers.AwsRegions, regionType string) string {
	str := ""

	if nil != regions {
		str += `
		  ` + regionType + `:`
		for _, region := range regions {
			str += `
		    - name: ` + region.Name
		}
	} else {
		str += `
		  ` + regionType + `: []`
	}

	return str
}

func getAwsLifecycleHooks(lifeCycles []*providers.LifecycleHooks) string {
	str := ""

	if nil != lifeCycles {
		str += `
		  lifecycleHooks:`
		for _, lifeCycle := range lifeCycles {
			str += `
		    - defaultResult: ` + lifeCycle.DefaultResult + `
		      heartbeatTimeout: ` + helpers.IntToString(lifeCycle.HeartbeatTimeout) + `
		      lifecycleTransition: ` + lifeCycle.LifecycleTransition + `
		      notificationTargetARN: ` + lifeCycle.NotificationTargetARN + `
		      roleARN: ` + lifeCycle.RoleARN
		}
	} else {
		str += `
		  lifecycleHooks: []`
	}

	return str
}

func GetAwsBakeryDefault(bakeryDefaults *providers.AwsBakeryDefaults) string {
	str := ""

	//TODO Check null pointer
	if nil != bakeryDefaults {
		str += `
		  bakeryDefaults:` + `
		    templateFile: ` + bakeryDefaults.TemplateFile + `
		    awsAccessKey: ` + bakeryDefaults.AwsAccessKey + `
		    awsSecretKey: ` + bakeryDefaults.AwsSecretKey +
			GetAwsBaseImages(bakeryDefaults.BaseImages) + `
		    awsAssociatePublicIpAddress: ` + strconv.FormatBool(bakeryDefaults.AwsAssociatePublicIpAddress) + `
		    defaultVirtualizationType: ` + bakeryDefaults.DefaultVirtualizationType + `
		    awsSubnetId: ` + bakeryDefaults.AwsSubnetId + `
		    awsVpcId: ` + bakeryDefaults.AwsVpcId
	} else {
		str += `
		  bakeryDefaults: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetAwsBaseImages(baseImages *providers.AWSBaseImages) string {
	str := ""

	if nil != baseImages {
		str += `
		    baseImages:`
		if nil != baseImages.BaseImage {
			for _, baseImage := range baseImages.BaseImage {
				str += `
		      - baseImage:
		        id: ` + baseImage.Id + `
		        shortDescription: ` + baseImage.ShortDescription + `
		        detailedDescription: ` + baseImage.DetailedDescription + `
		        packageType: ` + baseImage.PackageType + `
		        templateFile: ` + baseImage.TemplateFile
				if nil != baseImage.VirtualizationSettings {
					str += `
		        virtualizationSettings:`
					for _, virtualSettings := range baseImage.VirtualizationSettings {
						str += `
		          - region: ` + virtualSettings.Region + `
		            virtualizationType: ` + virtualSettings.VirtualizationType + `
		            instanceType: ` + virtualSettings.InstanceType + `
		            sourceAmi: ` + virtualSettings.SourceAmi + /*`
							  winRmUserName: ` + virtualSettings.WinRmUserName + */`
		            spotPrice: ` + virtualSettings.SpotPrice + `
		            spotPriceAutoProduct: ` + virtualSettings.SpotPriceAutoProduct
					}
				} else {
					str += `
		        virtualizationSettings: []`
				}
			}
		}
	} else {
		str += `
		    baseImages: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
