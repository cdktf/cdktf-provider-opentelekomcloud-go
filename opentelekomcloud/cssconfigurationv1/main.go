// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cssconfigurationv1

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.cssConfigurationV1.CssConfigurationV1",
		reflect.TypeOf((*CssConfigurationV1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "autoCreateIndex", GoGetter: "AutoCreateIndex"},
			_jsii_.MemberProperty{JsiiProperty: "autoCreateIndexInput", GoGetter: "AutoCreateIndexInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clusterId", GoGetter: "ClusterId"},
			_jsii_.MemberProperty{JsiiProperty: "clusterIdInput", GoGetter: "ClusterIdInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "httpCorsAllowCredentials", GoGetter: "HttpCorsAllowCredentials"},
			_jsii_.MemberProperty{JsiiProperty: "httpCorsAllowCredentialsInput", GoGetter: "HttpCorsAllowCredentialsInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpCorsAllowHeaders", GoGetter: "HttpCorsAllowHeaders"},
			_jsii_.MemberProperty{JsiiProperty: "httpCorsAllowHeadersInput", GoGetter: "HttpCorsAllowHeadersInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpCorsAllowMethods", GoGetter: "HttpCorsAllowMethods"},
			_jsii_.MemberProperty{JsiiProperty: "httpCorsAllowMethodsInput", GoGetter: "HttpCorsAllowMethodsInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpCorsAllowOrigin", GoGetter: "HttpCorsAllowOrigin"},
			_jsii_.MemberProperty{JsiiProperty: "httpCorsAllowOriginInput", GoGetter: "HttpCorsAllowOriginInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpCorsEnabled", GoGetter: "HttpCorsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "httpCorsEnabledInput", GoGetter: "HttpCorsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "httpCorsMaxAge", GoGetter: "HttpCorsMaxAge"},
			_jsii_.MemberProperty{JsiiProperty: "httpCorsMaxAgeInput", GoGetter: "HttpCorsMaxAgeInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberProperty{JsiiProperty: "indicesQueriesCacheSize", GoGetter: "IndicesQueriesCacheSize"},
			_jsii_.MemberProperty{JsiiProperty: "indicesQueriesCacheSizeInput", GoGetter: "IndicesQueriesCacheSizeInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "reindexRemoteWhitelist", GoGetter: "ReindexRemoteWhitelist"},
			_jsii_.MemberProperty{JsiiProperty: "reindexRemoteWhitelistInput", GoGetter: "ReindexRemoteWhitelistInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoCreateIndex", GoMethod: "ResetAutoCreateIndex"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpCorsAllowCredentials", GoMethod: "ResetHttpCorsAllowCredentials"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpCorsAllowHeaders", GoMethod: "ResetHttpCorsAllowHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpCorsAllowMethods", GoMethod: "ResetHttpCorsAllowMethods"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpCorsAllowOrigin", GoMethod: "ResetHttpCorsAllowOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpCorsEnabled", GoMethod: "ResetHttpCorsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetHttpCorsMaxAge", GoMethod: "ResetHttpCorsMaxAge"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIndicesQueriesCacheSize", GoMethod: "ResetIndicesQueriesCacheSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetReindexRemoteWhitelist", GoMethod: "ResetReindexRemoteWhitelist"},
			_jsii_.MemberMethod{JsiiMethod: "resetThreadPoolForceMergeSize", GoMethod: "ResetThreadPoolForceMergeSize"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "threadPoolForceMergeSize", GoGetter: "ThreadPoolForceMergeSize"},
			_jsii_.MemberProperty{JsiiProperty: "threadPoolForceMergeSizeInput", GoGetter: "ThreadPoolForceMergeSizeInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_CssConfigurationV1{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.cssConfigurationV1.CssConfigurationV1Config",
		reflect.TypeOf((*CssConfigurationV1Config)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.cssConfigurationV1.CssConfigurationV1Timeouts",
		reflect.TypeOf((*CssConfigurationV1Timeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.cssConfigurationV1.CssConfigurationV1TimeoutsOutputReference",
		reflect.TypeOf((*CssConfigurationV1TimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_CssConfigurationV1TimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
