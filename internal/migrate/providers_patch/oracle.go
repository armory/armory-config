package providers_patch

import (
	"strings"

	"github.com/austinthao5/golang_proto_test/config/deploymentConfigurations/providers"
	"github.com/austinthao5/golang_proto_test/internal/helpers"
)

func (ProvidersData *Providers) SetOracle(providersRef *providers.Providers) error {

	if nil != providersRef.Oracle {
		if providersRef.Oracle.Enabled {
			ProvidersData.Enable = "oracle"
		}

		str := helpers.PrintFmtBool(`enabled: `, providersRef.Oracle.Enabled, 5, true) +
			helpers.PrintFmtStr(`primaryAccount: `, providersRef.Oracle.PrimaryAccount, 5, true) +
			GetOracleAccounts(providersRef.Oracle.Accounts) +
			GetOracleBakeryDefaultsAccounts(providersRef.Oracle.BakeryDefaults)

		str = strings.Replace(str, "\t", "          ", -1)
		ProvidersData.Oracle = str
	} else {
		ProvidersData.Oracle = " {}"
		// 	return fmt.Errorf("Oracle value is null")
	}

	return nil
}

func GetOracleAccounts(accounts []*providers.OracleAccounts) string {
	str := ""

	if nil != accounts {
		str += `
		  accounts:`
		for _, account := range accounts {
			str += helpers.PrintFmtStr(`- name: `, account.Name, 6, true) +
				helpers.PrintFmtStr(`environment: `, account.Environment, 7, true) +
				getProvidersStringArrayAppend(account.RequiredGroupMembership, "requiredGroupMembership", "- ") +
				strings.Replace(getPermissions(account.Permissions), "\t", "     ", -1) +
				helpers.PrintFmtStr(`compartmentId: `, account.CompartmentId, 7, true) +
				helpers.PrintFmtStr(`userId: `, account.UserId, 7, true) +
				helpers.PrintFmtStr(`fingerprint: `, account.Fingerprint, 7, true) +
				helpers.PrintFmtStr(`sshPrivateKeyFilePath: `, account.SshPrivateKeyFilePath, 7, true) +
				helpers.PrintFmtStr(`privateKeyPassphrase: `, account.PrivateKeyPassphrase, 7, true) +
				helpers.PrintFmtStr(`tenancyId: `, account.TenancyId, 7, true) +
				helpers.PrintFmtStr(`region: `, account.Region, 7, true)
			//   getProvidersStringArray(account.Regions, "regions") + `
		}
	} else {
		str += `
		  accounts: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetOracleBakeryDefaultsAccounts(bakeryDefault *providers.OracleBakery) string {
	str := ""

	if nil != bakeryDefault {
		str += `
		  bakeryDefaults:` +
			helpers.PrintFmtStr(`templateFile: `, bakeryDefault.TemplateFile, 6, true) +
			GetOracleBaseImages(bakeryDefault.BaseImages) +
			helpers.PrintFmtStr(`availabilityDomain: `, bakeryDefault.AvailabilityDomain, 6, true) +
			helpers.PrintFmtStr(`subnetId: `, bakeryDefault.SubnetId, 6, true) +
			helpers.PrintFmtStr(`instanceShape: `, bakeryDefault.InstanceShape, 6, true)
	} else {
		str += `
		  bakeryDefaults: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}

func GetOracleBaseImages(baseImages []*providers.OracleBaseImages) string {
	str := ""

	if nil != baseImages {
		str += `
		    baseImages:`
		for _, baseImage := range baseImages {
			if nil != baseImage.BaseImage {
				str += `
			  - baseImage:` +
					helpers.PrintFmtStr(`id: `, baseImage.BaseImage.Id, 8, true) +
					helpers.PrintFmtStr(`shortDescription: `, baseImage.BaseImage.ShortDescription, 8, true) +
					helpers.PrintFmtStr(`detailedDescription: `, baseImage.BaseImage.DetailedDescription, 8, true) +
					helpers.PrintFmtStr(`packageType: `, baseImage.BaseImage.PackageType, 8, true) +
					helpers.PrintFmtStr(`templateFile: `, baseImage.BaseImage.TemplateFile, 8, true)
			}
			//for _, virtualSetting := range baseImage.VirtualizationSettings {
			if nil != baseImage.VirtualizationSettings {
				str += `
			    virtualizationSettings:` +
					helpers.PrintFmtStr(`baseImageId: `, baseImage.VirtualizationSettings.BaseImageId, 9, true) +
					helpers.PrintFmtStr(`sshUserName: `, baseImage.VirtualizationSettings.SshUserName, 9, true)
			}
			//}
		}
	} else {
		str += `
		    baseImages: []`
	}

	str = strings.Replace(str, "\t", "    ", -1)
	return str
}
