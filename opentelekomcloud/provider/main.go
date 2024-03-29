// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.provider.OpentelekomcloudProvider",
		reflect.TypeOf((*OpentelekomcloudProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessKey", GoGetter: "AccessKey"},
			_jsii_.MemberProperty{JsiiProperty: "accessKeyInput", GoGetter: "AccessKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "agencyDomainName", GoGetter: "AgencyDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "agencyDomainNameInput", GoGetter: "AgencyDomainNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "agencyName", GoGetter: "AgencyName"},
			_jsii_.MemberProperty{JsiiProperty: "agencyNameInput", GoGetter: "AgencyNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowReauth", GoGetter: "AllowReauth"},
			_jsii_.MemberProperty{JsiiProperty: "allowReauthInput", GoGetter: "AllowReauthInput"},
			_jsii_.MemberProperty{JsiiProperty: "authUrl", GoGetter: "AuthUrl"},
			_jsii_.MemberProperty{JsiiProperty: "authUrlInput", GoGetter: "AuthUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "backoffRetryTimeout", GoGetter: "BackoffRetryTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "backoffRetryTimeoutInput", GoGetter: "BackoffRetryTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "cacertFile", GoGetter: "CacertFile"},
			_jsii_.MemberProperty{JsiiProperty: "cacertFileInput", GoGetter: "CacertFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "cert", GoGetter: "Cert"},
			_jsii_.MemberProperty{JsiiProperty: "certInput", GoGetter: "CertInput"},
			_jsii_.MemberProperty{JsiiProperty: "cloud", GoGetter: "Cloud"},
			_jsii_.MemberProperty{JsiiProperty: "cloudInput", GoGetter: "CloudInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "delegatedProject", GoGetter: "DelegatedProject"},
			_jsii_.MemberProperty{JsiiProperty: "delegatedProjectInput", GoGetter: "DelegatedProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "domainId", GoGetter: "DomainId"},
			_jsii_.MemberProperty{JsiiProperty: "domainIdInput", GoGetter: "DomainIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "domainName", GoGetter: "DomainName"},
			_jsii_.MemberProperty{JsiiProperty: "domainNameInput", GoGetter: "DomainNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "endpointType", GoGetter: "EndpointType"},
			_jsii_.MemberProperty{JsiiProperty: "endpointTypeInput", GoGetter: "EndpointTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "insecure", GoGetter: "Insecure"},
			_jsii_.MemberProperty{JsiiProperty: "insecureInput", GoGetter: "InsecureInput"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxBackoffRetries", GoGetter: "MaxBackoffRetries"},
			_jsii_.MemberProperty{JsiiProperty: "maxBackoffRetriesInput", GoGetter: "MaxBackoffRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetries", GoGetter: "MaxRetries"},
			_jsii_.MemberProperty{JsiiProperty: "maxRetriesInput", GoGetter: "MaxRetriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passcode", GoGetter: "Passcode"},
			_jsii_.MemberProperty{JsiiProperty: "passcodeInput", GoGetter: "PasscodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "password", GoGetter: "Password"},
			_jsii_.MemberProperty{JsiiProperty: "passwordInput", GoGetter: "PasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessKey", GoMethod: "ResetAccessKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetAgencyDomainName", GoMethod: "ResetAgencyDomainName"},
			_jsii_.MemberMethod{JsiiMethod: "resetAgencyName", GoMethod: "ResetAgencyName"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowReauth", GoMethod: "ResetAllowReauth"},
			_jsii_.MemberMethod{JsiiMethod: "resetAuthUrl", GoMethod: "ResetAuthUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackoffRetryTimeout", GoMethod: "ResetBackoffRetryTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetCacertFile", GoMethod: "ResetCacertFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetCert", GoMethod: "ResetCert"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloud", GoMethod: "ResetCloud"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelegatedProject", GoMethod: "ResetDelegatedProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomainId", GoMethod: "ResetDomainId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDomainName", GoMethod: "ResetDomainName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndpointType", GoMethod: "ResetEndpointType"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecure", GoMethod: "ResetInsecure"},
			_jsii_.MemberMethod{JsiiMethod: "resetKey", GoMethod: "ResetKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxBackoffRetries", GoMethod: "ResetMaxBackoffRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRetries", GoMethod: "ResetMaxRetries"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasscode", GoMethod: "ResetPasscode"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassword", GoMethod: "ResetPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretKey", GoMethod: "ResetSecretKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityToken", GoMethod: "ResetSecurityToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetSwauth", GoMethod: "ResetSwauth"},
			_jsii_.MemberMethod{JsiiMethod: "resetTenantId", GoMethod: "ResetTenantId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTenantName", GoMethod: "ResetTenantName"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserId", GoMethod: "ResetUserId"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserName", GoMethod: "ResetUserName"},
			_jsii_.MemberProperty{JsiiProperty: "secretKey", GoGetter: "SecretKey"},
			_jsii_.MemberProperty{JsiiProperty: "secretKeyInput", GoGetter: "SecretKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityToken", GoGetter: "SecurityToken"},
			_jsii_.MemberProperty{JsiiProperty: "securityTokenInput", GoGetter: "SecurityTokenInput"},
			_jsii_.MemberProperty{JsiiProperty: "swauth", GoGetter: "Swauth"},
			_jsii_.MemberProperty{JsiiProperty: "swauthInput", GoGetter: "SwauthInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tenantId", GoGetter: "TenantId"},
			_jsii_.MemberProperty{JsiiProperty: "tenantIdInput", GoGetter: "TenantIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "tenantName", GoGetter: "TenantName"},
			_jsii_.MemberProperty{JsiiProperty: "tenantNameInput", GoGetter: "TenantNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userId", GoGetter: "UserId"},
			_jsii_.MemberProperty{JsiiProperty: "userIdInput", GoGetter: "UserIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "userName", GoGetter: "UserName"},
			_jsii_.MemberProperty{JsiiProperty: "userNameInput", GoGetter: "UserNameInput"},
		},
		func() interface{} {
			j := jsiiProxy_OpentelekomcloudProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.provider.OpentelekomcloudProviderConfig",
		reflect.TypeOf((*OpentelekomcloudProviderConfig)(nil)).Elem(),
	)
}
