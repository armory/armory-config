package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetAwsData(providersRef *providers.Providers) error {

	if nil != providersRef.Aws {
		//To avoid null pointers
		featuresCloudF := ""
		featuresLambda := ""
		if nil != providersRef.Aws.Features {
			if nil != providersRef.Aws.Features.CloudFormation {
				featuresCloudF = strconv.FormatBool(providersRef.Aws.Features.CloudFormation.Enabled)
			}
			if nil != providersRef.Aws.Features.Lambda {
				featuresLambda = strconv.FormatBool(providersRef.Aws.Features.Lambda.Enabled)
			}
		}

		if providersRef.Aws.Enabled {
			ProvidersData.Enable = "aws"
		}

		str := helpers.PrintFmtBool(`enabled: `, providersRef.Aws.Enabled, 5, true) +
			helpers.PrintFmtStr(`primaryAccount: `, providersRef.Aws.PrimaryAccount, 5, true) +
			GetAwsAccounts(providersRef) +
			GetAwsBakeryDefault(providersRef.Aws.BakeryDefaults) +
			helpers.PrintFmtStrApostrophe(`accessKeyId: `, providersRef.Aws.AccessKeyId, 5, true) +
			helpers.PrintFmtStrApostrophe(`secretAccessKey: `, providersRef.Aws.SecretAccessKey, 5, true) +
			helpers.PrintFmtStrApostrophe(`defaultKeyPairTemplate: `, providersRef.Aws.DefaultKeyPairTemplate, 5, true) +
			strings.Replace(getAwsRegions(providersRef.Aws.DefaultRegions, "defaultRegions"), "\t", "    ", -1) + `
	defaults:` +
			helpers.PrintFmtStr(`iamRole: `, providersRef.Aws.Defaults.IamRole, 6, true)

		if featuresLambda == `true` || featuresCloudF == `true` {

			str += `
	features:`

			if featuresLambda == `true` {
				str += `
  	lambda:` +
					helpers.PrintFmtStr(`enabled: `, featuresLambda, 7, true)
			}
			if featuresCloudF == `true` {
				str += `
  	cloudFormation:` +
					helpers.PrintFmtStr(`enabled: `, featuresCloudF, 7, true)
			}

		}
		// features:
		//   cloudFormation:` +
		// 		helpers.PrintFmtStr(`enabled: `, featuresCloudF, 7, true) + `
		//   lambda:` +
		// 		helpers.PrintFmtStr(`enabled: `, featuresLambda, 7, true)

		str = strings.Replace(str, "\t", "          ", -1)

		ProvidersData.Aws = str
	} else {
		ProvidersData.Aws = " {}"
		// 	return fmt.Errorf("Aws value is null")
	}

	return nil
}

func GetAwsAccounts(provider *providers.Providers) string {
	str := ""

	if nil != provider.Aws.Accounts {
		str += `
		  accounts:`
		for _, account := range provider.Aws.Accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 5, true) +
				helpers.PrintFmtStr(`environment: `, account.Environment, 6, true) +
				strings.Replace(getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- "), "\t", "   ", -1) +
				getPermissions(account.Permissions) +
				helpers.PrintFmtStr(`defaultKeyPair: `, account.DefaultKeyPair, 6, true) +
				helpers.PrintFmtStr(`edda: `, account.Edda, 6, true) +
				helpers.PrintFmtStr(`discovery: `, account.Discovery, 6, true) +
				helpers.PrintFmtStrApostrophe(`accountId: `, account.AccountId, 6, true) +
				strings.Replace(getAwsRegions(account.Regions, "regions"), "\t", "     ", -1) +
				strings.Replace(getAwsLifecycleHooks(account.LifecycleHooks), "\t", "     ", -1) +
				helpers.PrintFmtStrApostrophe(`assumeRole: `, account.AssumeRole, 6, true)
			//TODO Check if providerVersion Missing proto
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
			str += helpers.PrintFmtStr(`- defaultResult: `, lifeCycle.DefaultResult, 7, true) +
				helpers.PrintFmtInt(`heartbeatTimeout: `, lifeCycle.HeartbeatTimeout, 8, true) +
				helpers.PrintFmtStr(`lifecycleTransition: `, lifeCycle.LifecycleTransition, 8, true) +
				helpers.PrintFmtStr(`notificationTargetARN: `, lifeCycle.NotificationTargetARN, 8, true) +
				helpers.PrintFmtStr(`roleARN: `, lifeCycle.RoleARN, 8, true)
		}
	} else {
		str += `
		  lifecycleHooks: []`
	}

	return str
}

func GetAwsBakeryDefault(bakeryDefaults *providers.AwsBakeryDefaults) string {
	str := ""

	if nil != bakeryDefaults {
		str += `
		  bakeryDefaults:` +
			helpers.PrintFmtStr(`templateFile: `, bakeryDefaults.TemplateFile, 6, true) +
			helpers.PrintFmtStrApostrophe(`awsAccessKey: `, bakeryDefaults.AwsAccessKey, 6, true) +
			helpers.PrintFmtStr(`awsSecretKey: `, bakeryDefaults.AwsSecretKey, 6, true) +
			GetAwsBaseImages(bakeryDefaults.BaseImages) +
			helpers.PrintFmtBool(`awsAssociatePublicIpAddress: `, bakeryDefaults.AwsAssociatePublicIpAddress, 6, true) +
			helpers.PrintFmtStr(`defaultVirtualizationType: `, bakeryDefaults.DefaultVirtualizationType, 6, true) +
			helpers.PrintFmtStr(`awsSubnetId: `, bakeryDefaults.AwsSubnetId, 6, true) +
			helpers.PrintFmtStr(`awsVpcId: `, bakeryDefaults.AwsVpcId, 6, true)
	} else {
		str += `
		  bakeryDefaults: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetAwsBaseImages(baseImages []*providers.AWSBaseImages) string {
	str := ""

	if nil != baseImages {
		str += `
		    baseImages:`
		for _, baseImageArray := range baseImages {
			if nil != baseImageArray.BaseImage {
				baseImage := baseImageArray.BaseImage
				str += `
		      - baseImage:` +
					helpers.PrintFmtStr(`id: `, baseImage.Id, 9, true) +
					helpers.PrintFmtStr(`shortDescription: `, baseImage.ShortDescription, 9, true) +
					helpers.PrintFmtStr(`detailedDescription: `, baseImage.DetailedDescription, 9, true) +
					helpers.PrintFmtStr(`packageType: `, baseImage.PackageType, 9, true) +
					helpers.PrintFmtStr(`templateFile: `, baseImage.TemplateFile, 9, true)
			}
			if nil != baseImageArray.VirtualizationSettings {
				str += `
		        virtualizationSettings:`
				for _, virtualSettings := range baseImageArray.VirtualizationSettings {
					str += helpers.PrintFmtStr(`- region: `, virtualSettings.Region, 9, true) +
						helpers.PrintFmtStr(`virtualizationType: `, virtualSettings.VirtualizationType, 10, true) +
						helpers.PrintFmtStr(`instanceType: `, virtualSettings.InstanceType, 10, true) +
						helpers.PrintFmtStr(`sourceAmi: `, virtualSettings.SourceAmi, 10, true) +
						helpers.PrintFmtStr(`winRmUserName: `, virtualSettings.WinRmUserName, 10, true) +
						helpers.PrintFmtStr(`spotPrice: `, virtualSettings.SpotPrice, 10, true) +
						helpers.PrintFmtStr(`spotPriceAutoProduct: `, virtualSettings.SpotPriceAutoProduct, 10, true)
				}
			} else {
				str += `
		        virtualizationSettings: []`
			}
		}
	} else {
		str += `
		    baseImages: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
