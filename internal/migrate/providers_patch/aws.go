package providers_patch

import (
	"strconv"
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
)

func (ProvidersData *Providers) SetAwsData(providersRef *providers.Providers) error {

	// if nil != providersRef.Aws {
	// 	return fmt.Errorf("Aws value is null")
	// }

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
	  iamRole: ` + providersRef.Aws.Defaults.IamRole /*+ `
	features:
	  cloudFormation:
	    enabled: ` + strconv.FormatBool(providersRef.Aws.Features.Cloudformation.Enabled) + `                       # (Default: false). Enable cloudformation support on this AWS account.` + `
	  lambda:
	    enabled: ` + strconv.FormatBool(providersRef.Aws.Features.Lambda.Enabled) */
	//TODO Fix this

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
				strings.Replace(getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- "), "\t", "   ", -1) + `
		    defaultKeyPair: ` + account.DefaultKeyPair + `
		    edda : ` + account.Edda + `
		    discovery: ` + account.Discovery + `
		    accountId: ` + account.AccountId +
				strings.Replace(getAwsRegions(account.Regions, "region"), "\t", "     ", -1) +
				strings.Replace(getAwsLifecycleHooks(account.LifecycleHooks), "\t", "     ", -1) + `
		    Permission: {}` //TODO + account.Permission`
			//TODO assumeRole
			//TODO providerVersion
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
		      heartbeatTimeout: ` + strconv.FormatInt(int64(lifeCycle.HeartbeatTimeout), 10) + `
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

// TODO
func GetAwsBakeryDefault(bakeryDefaults *providers.AwsBakeryDefaults) string {
	str := ""

	if nil != bakeryDefaults {
		str += `
		  bakeryDefaults:                        # Configuration for Spinnaker’s image bakery.Configuration for Spinnaker’s image bakery.`
	} else {
		str += `
		  bakeryDefaults: []                         # Configuration for Spinnaker’s image bakery.Configuration for Spinnaker’s image bakery.`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
