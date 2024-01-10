// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudcceaddontemplatev3

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCceAddonTemplateV3.DataOpentelekomcloudCceAddonTemplateV3",
		reflect.TypeOf((*DataOpentelekomcloudCceAddonTemplateV3)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "addonName", GoGetter: "AddonName"},
			_jsii_.MemberProperty{JsiiProperty: "addonNameInput", GoGetter: "AddonNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "addonVersion", GoGetter: "AddonVersion"},
			_jsii_.MemberProperty{JsiiProperty: "addonVersionInput", GoGetter: "AddonVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIp", GoGetter: "ClusterIp"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIpInput", GoGetter: "ClusterIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "clusterVersions", GoGetter: "ClusterVersions"},
			_jsii_.MemberProperty{JsiiProperty: "clusterVersionsInput", GoGetter: "ClusterVersionsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersion", GoGetter: "ImageVersion"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionInput", GoGetter: "ImageVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetClusterIp", GoMethod: "ResetClusterIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetClusterVersions", GoMethod: "ResetClusterVersions"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageVersion", GoMethod: "ResetImageVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSwrAddr", GoMethod: "ResetSwrAddr"},
			_jsii_.MemberMethod{JsiiMethod: "resetSwrUser", GoMethod: "ResetSwrUser"},
			_jsii_.MemberProperty{JsiiProperty: "swrAddr", GoGetter: "SwrAddr"},
			_jsii_.MemberProperty{JsiiProperty: "swrAddrInput", GoGetter: "SwrAddrInput"},
			_jsii_.MemberProperty{JsiiProperty: "swrUser", GoGetter: "SwrUser"},
			_jsii_.MemberProperty{JsiiProperty: "swrUserInput", GoGetter: "SwrUserInput"},
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
			j := jsiiProxy_DataOpentelekomcloudCceAddonTemplateV3{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudCceAddonTemplateV3.DataOpentelekomcloudCceAddonTemplateV3Config",
		reflect.TypeOf((*DataOpentelekomcloudCceAddonTemplateV3Config)(nil)).Elem(),
	)
}
