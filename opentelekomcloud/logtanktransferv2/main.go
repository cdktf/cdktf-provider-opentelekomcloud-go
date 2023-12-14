// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logtanktransferv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.logtankTransferV2.LogtankTransferV2",
		reflect.TypeOf((*LogtankTransferV2)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createTime", GoGetter: "CreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "dirPrefixName", GoGetter: "DirPrefixName"},
			_jsii_.MemberProperty{JsiiProperty: "dirPrefixNameInput", GoGetter: "DirPrefixNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupId", GoGetter: "LogGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupIdInput", GoGetter: "LogGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamIds", GoGetter: "LogStreamIds"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamIdsInput", GoGetter: "LogStreamIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "logTransferMode", GoGetter: "LogTransferMode"},
			_jsii_.MemberProperty{JsiiProperty: "logTransferType", GoGetter: "LogTransferType"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "obsBucketName", GoGetter: "ObsBucketName"},
			_jsii_.MemberProperty{JsiiProperty: "obsBucketNameInput", GoGetter: "ObsBucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "obsEncryptionEnable", GoGetter: "ObsEncryptionEnable"},
			_jsii_.MemberProperty{JsiiProperty: "obsEncryptionId", GoGetter: "ObsEncryptionId"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "period", GoGetter: "Period"},
			_jsii_.MemberProperty{JsiiProperty: "periodInput", GoGetter: "PeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "periodUnit", GoGetter: "PeriodUnit"},
			_jsii_.MemberProperty{JsiiProperty: "periodUnitInput", GoGetter: "PeriodUnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "prefixName", GoGetter: "PrefixName"},
			_jsii_.MemberProperty{JsiiProperty: "prefixNameInput", GoGetter: "PrefixNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDirPrefixName", GoMethod: "ResetDirPrefixName"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrefixName", GoMethod: "ResetPrefixName"},
			_jsii_.MemberMethod{JsiiMethod: "resetSwitchOn", GoMethod: "ResetSwitchOn"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "storageFormat", GoGetter: "StorageFormat"},
			_jsii_.MemberProperty{JsiiProperty: "storageFormatInput", GoGetter: "StorageFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "switchOn", GoGetter: "SwitchOn"},
			_jsii_.MemberProperty{JsiiProperty: "switchOnInput", GoGetter: "SwitchOnInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_LogtankTransferV2{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.logtankTransferV2.LogtankTransferV2Config",
		reflect.TypeOf((*LogtankTransferV2Config)(nil)).Elem(),
	)
}
