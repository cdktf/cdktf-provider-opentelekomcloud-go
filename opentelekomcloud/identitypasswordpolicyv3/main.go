// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitypasswordpolicyv3

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.identityPasswordPolicyV3.IdentityPasswordPolicyV3",
		reflect.TypeOf((*IdentityPasswordPolicyV3)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "maximumConsecutiveIdenticalChars", GoGetter: "MaximumConsecutiveIdenticalChars"},
			_jsii_.MemberProperty{JsiiProperty: "maximumConsecutiveIdenticalCharsInput", GoGetter: "MaximumConsecutiveIdenticalCharsInput"},
			_jsii_.MemberProperty{JsiiProperty: "maximumPasswordLength", GoGetter: "MaximumPasswordLength"},
			_jsii_.MemberProperty{JsiiProperty: "minimumPasswordAge", GoGetter: "MinimumPasswordAge"},
			_jsii_.MemberProperty{JsiiProperty: "minimumPasswordAgeInput", GoGetter: "MinimumPasswordAgeInput"},
			_jsii_.MemberProperty{JsiiProperty: "minimumPasswordLength", GoGetter: "MinimumPasswordLength"},
			_jsii_.MemberProperty{JsiiProperty: "minimumPasswordLengthInput", GoGetter: "MinimumPasswordLengthInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "numberOfRecentPasswordsDisallowed", GoGetter: "NumberOfRecentPasswordsDisallowed"},
			_jsii_.MemberProperty{JsiiProperty: "numberOfRecentPasswordsDisallowedInput", GoGetter: "NumberOfRecentPasswordsDisallowedInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passwordCharCombination", GoGetter: "PasswordCharCombination"},
			_jsii_.MemberProperty{JsiiProperty: "passwordCharCombinationInput", GoGetter: "PasswordCharCombinationInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordNotUsernameOrInvert", GoGetter: "PasswordNotUsernameOrInvert"},
			_jsii_.MemberProperty{JsiiProperty: "passwordNotUsernameOrInvertInput", GoGetter: "PasswordNotUsernameOrInvertInput"},
			_jsii_.MemberProperty{JsiiProperty: "passwordRequirements", GoGetter: "PasswordRequirements"},
			_jsii_.MemberProperty{JsiiProperty: "passwordValidityPeriod", GoGetter: "PasswordValidityPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "passwordValidityPeriodInput", GoGetter: "PasswordValidityPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaximumConsecutiveIdenticalChars", GoMethod: "ResetMaximumConsecutiveIdenticalChars"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimumPasswordAge", GoMethod: "ResetMinimumPasswordAge"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinimumPasswordLength", GoMethod: "ResetMinimumPasswordLength"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumberOfRecentPasswordsDisallowed", GoMethod: "ResetNumberOfRecentPasswordsDisallowed"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordCharCombination", GoMethod: "ResetPasswordCharCombination"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordNotUsernameOrInvert", GoMethod: "ResetPasswordNotUsernameOrInvert"},
			_jsii_.MemberMethod{JsiiMethod: "resetPasswordValidityPeriod", GoMethod: "ResetPasswordValidityPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_IdentityPasswordPolicyV3{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.identityPasswordPolicyV3.IdentityPasswordPolicyV3Config",
		reflect.TypeOf((*IdentityPasswordPolicyV3Config)(nil)).Elem(),
	)
}
