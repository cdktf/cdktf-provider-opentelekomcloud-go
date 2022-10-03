package dataopentelekomcloudlblistenerv3

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudLbListenerV3.DataOpentelekomcloudLbListenerV3",
		reflect.TypeOf((*DataOpentelekomcloudLbListenerV3)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "adminStateUp", GoGetter: "AdminStateUp"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "clientCaTlsContainerRef", GoGetter: "ClientCaTlsContainerRef"},
			_jsii_.MemberProperty{JsiiProperty: "clientCaTlsContainerRefInput", GoGetter: "ClientCaTlsContainerRefInput"},
			_jsii_.MemberProperty{JsiiProperty: "clientTimeout", GoGetter: "ClientTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "clientTimeoutInput", GoGetter: "ClientTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createdAt", GoGetter: "CreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPoolId", GoGetter: "DefaultPoolId"},
			_jsii_.MemberProperty{JsiiProperty: "defaultPoolIdInput", GoGetter: "DefaultPoolIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTlsContainerRef", GoGetter: "DefaultTlsContainerRef"},
			_jsii_.MemberProperty{JsiiProperty: "defaultTlsContainerRefInput", GoGetter: "DefaultTlsContainerRefInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "http2Enable", GoGetter: "Http2Enable"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "insertHeaders", GoGetter: "InsertHeaders"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keepAliveTimeout", GoGetter: "KeepAliveTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "keepAliveTimeoutInput", GoGetter: "KeepAliveTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "loadbalancerId", GoGetter: "LoadbalancerId"},
			_jsii_.MemberProperty{JsiiProperty: "loadbalancerIdInput", GoGetter: "LoadbalancerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "memberAddress", GoGetter: "MemberAddress"},
			_jsii_.MemberProperty{JsiiProperty: "memberAddressInput", GoGetter: "MemberAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "memberDeviceId", GoGetter: "MemberDeviceId"},
			_jsii_.MemberProperty{JsiiProperty: "memberDeviceIdInput", GoGetter: "MemberDeviceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "memberRetryEnable", GoGetter: "MemberRetryEnable"},
			_jsii_.MemberProperty{JsiiProperty: "memberTimeout", GoGetter: "MemberTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "memberTimeoutInput", GoGetter: "MemberTimeoutInput"},
			_jsii_.MemberProperty{JsiiProperty: "memoryRetryEnable", GoGetter: "MemoryRetryEnable"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "protocol", GoGetter: "Protocol"},
			_jsii_.MemberProperty{JsiiProperty: "protocolInput", GoGetter: "ProtocolInput"},
			_jsii_.MemberProperty{JsiiProperty: "protocolPort", GoGetter: "ProtocolPort"},
			_jsii_.MemberProperty{JsiiProperty: "protocolPortInput", GoGetter: "ProtocolPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientCaTlsContainerRef", GoMethod: "ResetClientCaTlsContainerRef"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientTimeout", GoMethod: "ResetClientTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultPoolId", GoMethod: "ResetDefaultPoolId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultTlsContainerRef", GoMethod: "ResetDefaultTlsContainerRef"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeepAliveTimeout", GoMethod: "ResetKeepAliveTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoadbalancerId", GoMethod: "ResetLoadbalancerId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemberAddress", GoMethod: "ResetMemberAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemberDeviceId", GoMethod: "ResetMemberDeviceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetMemberTimeout", GoMethod: "ResetMemberTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProtocol", GoMethod: "ResetProtocol"},
			_jsii_.MemberMethod{JsiiMethod: "resetProtocolPort", GoMethod: "ResetProtocolPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetTlsCiphersPolicy", GoMethod: "ResetTlsCiphersPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "sniContainerRefs", GoGetter: "SniContainerRefs"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "tlsCiphersPolicy", GoGetter: "TlsCiphersPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "tlsCiphersPolicyInput", GoGetter: "TlsCiphersPolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "updatedAt", GoGetter: "UpdatedAt"},
		},
		func() interface{} {
			j := jsiiProxy_DataOpentelekomcloudLbListenerV3{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudLbListenerV3.DataOpentelekomcloudLbListenerV3Config",
		reflect.TypeOf((*DataOpentelekomcloudLbListenerV3Config)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudLbListenerV3.DataOpentelekomcloudLbListenerV3InsertHeaders",
		reflect.TypeOf((*DataOpentelekomcloudLbListenerV3InsertHeaders)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudLbListenerV3.DataOpentelekomcloudLbListenerV3InsertHeadersList",
		reflect.TypeOf((*DataOpentelekomcloudLbListenerV3InsertHeadersList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataOpentelekomcloudLbListenerV3InsertHeadersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudLbListenerV3.DataOpentelekomcloudLbListenerV3InsertHeadersOutputReference",
		reflect.TypeOf((*DataOpentelekomcloudLbListenerV3InsertHeadersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "forwardedForPort", GoGetter: "ForwardedForPort"},
			_jsii_.MemberProperty{JsiiProperty: "forwardedHost", GoGetter: "ForwardedHost"},
			_jsii_.MemberProperty{JsiiProperty: "forwardedPort", GoGetter: "ForwardedPort"},
			_jsii_.MemberProperty{JsiiProperty: "forwardElbIp", GoGetter: "ForwardElbIp"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataOpentelekomcloudLbListenerV3InsertHeadersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
