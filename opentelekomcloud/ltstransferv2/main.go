// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltstransferv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2",
		reflect.TypeOf((*LtsTransferV2)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "logGroupId", GoGetter: "LogGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupIdInput", GoGetter: "LogGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "logGroupName", GoGetter: "LogGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreams", GoGetter: "LogStreams"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamsInput", GoGetter: "LogStreamsInput"},
			_jsii_.MemberProperty{JsiiProperty: "logTransferInfo", GoGetter: "LogTransferInfo"},
			_jsii_.MemberProperty{JsiiProperty: "logTransferInfoInput", GoGetter: "LogTransferInfoInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putLogStreams", GoMethod: "PutLogStreams"},
			_jsii_.MemberMethod{JsiiMethod: "putLogTransferInfo", GoMethod: "PutLogTransferInfo"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
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
			j := jsiiProxy_LtsTransferV2{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2Config",
		reflect.TypeOf((*LtsTransferV2Config)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogStreams",
		reflect.TypeOf((*LtsTransferV2LogStreams)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogStreamsList",
		reflect.TypeOf((*LtsTransferV2LogStreamsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_LtsTransferV2LogStreamsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogStreamsOutputReference",
		reflect.TypeOf((*LtsTransferV2LogStreamsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "logStreamId", GoGetter: "LogStreamId"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamIdInput", GoGetter: "LogStreamIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamName", GoGetter: "LogStreamName"},
			_jsii_.MemberProperty{JsiiProperty: "logStreamNameInput", GoGetter: "LogStreamNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogStreamName", GoMethod: "ResetLogStreamName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LtsTransferV2LogStreamsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogTransferInfo",
		reflect.TypeOf((*LtsTransferV2LogTransferInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogTransferInfoLogAgencyTransfer",
		reflect.TypeOf((*LtsTransferV2LogTransferInfoLogAgencyTransfer)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference",
		reflect.TypeOf((*LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agencyDomainId", GoGetter: "AgencyDomainId"},
			_jsii_.MemberProperty{JsiiProperty: "agencyDomainIdInput", GoGetter: "AgencyDomainIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "agencyDomainName", GoGetter: "AgencyDomainName"},
			_jsii_.MemberProperty{JsiiProperty: "agencyDomainNameInput", GoGetter: "AgencyDomainNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "agencyName", GoGetter: "AgencyName"},
			_jsii_.MemberProperty{JsiiProperty: "agencyNameInput", GoGetter: "AgencyNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "agencyProjectId", GoGetter: "AgencyProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "agencyProjectIdInput", GoGetter: "AgencyProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			j := jsiiProxy_LtsTransferV2LogTransferInfoLogAgencyTransferOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogTransferInfoLogTransferDetail",
		reflect.TypeOf((*LtsTransferV2LogTransferInfoLogTransferDetail)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogTransferInfoLogTransferDetailOutputReference",
		reflect.TypeOf((*LtsTransferV2LogTransferInfoLogTransferDetailOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "obsBucketName", GoGetter: "ObsBucketName"},
			_jsii_.MemberProperty{JsiiProperty: "obsBucketNameInput", GoGetter: "ObsBucketNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "obsDirPrefixName", GoGetter: "ObsDirPrefixName"},
			_jsii_.MemberProperty{JsiiProperty: "obsDirPrefixNameInput", GoGetter: "ObsDirPrefixNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "obsEncryptedEnable", GoGetter: "ObsEncryptedEnable"},
			_jsii_.MemberProperty{JsiiProperty: "obsEncryptedEnableInput", GoGetter: "ObsEncryptedEnableInput"},
			_jsii_.MemberProperty{JsiiProperty: "obsEncryptedId", GoGetter: "ObsEncryptedId"},
			_jsii_.MemberProperty{JsiiProperty: "obsEncryptedIdInput", GoGetter: "ObsEncryptedIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "obsEpsId", GoGetter: "ObsEpsId"},
			_jsii_.MemberProperty{JsiiProperty: "obsEpsIdInput", GoGetter: "ObsEpsIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "obsPeriod", GoGetter: "ObsPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "obsPeriodInput", GoGetter: "ObsPeriodInput"},
			_jsii_.MemberProperty{JsiiProperty: "obsPeriodUnit", GoGetter: "ObsPeriodUnit"},
			_jsii_.MemberProperty{JsiiProperty: "obsPeriodUnitInput", GoGetter: "ObsPeriodUnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "obsPrefixName", GoGetter: "ObsPrefixName"},
			_jsii_.MemberProperty{JsiiProperty: "obsPrefixNameInput", GoGetter: "ObsPrefixNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "obsTimeZone", GoGetter: "ObsTimeZone"},
			_jsii_.MemberProperty{JsiiProperty: "obsTimeZoneId", GoGetter: "ObsTimeZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "obsTimeZoneIdInput", GoGetter: "ObsTimeZoneIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "obsTimeZoneInput", GoGetter: "ObsTimeZoneInput"},
			_jsii_.MemberProperty{JsiiProperty: "obsTransferPath", GoGetter: "ObsTransferPath"},
			_jsii_.MemberProperty{JsiiProperty: "obsTransferPathInput", GoGetter: "ObsTransferPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetObsBucketName", GoMethod: "ResetObsBucketName"},
			_jsii_.MemberMethod{JsiiMethod: "resetObsDirPrefixName", GoMethod: "ResetObsDirPrefixName"},
			_jsii_.MemberMethod{JsiiMethod: "resetObsEncryptedEnable", GoMethod: "ResetObsEncryptedEnable"},
			_jsii_.MemberMethod{JsiiMethod: "resetObsEncryptedId", GoMethod: "ResetObsEncryptedId"},
			_jsii_.MemberMethod{JsiiMethod: "resetObsEpsId", GoMethod: "ResetObsEpsId"},
			_jsii_.MemberMethod{JsiiMethod: "resetObsPeriod", GoMethod: "ResetObsPeriod"},
			_jsii_.MemberMethod{JsiiMethod: "resetObsPeriodUnit", GoMethod: "ResetObsPeriodUnit"},
			_jsii_.MemberMethod{JsiiMethod: "resetObsPrefixName", GoMethod: "ResetObsPrefixName"},
			_jsii_.MemberMethod{JsiiMethod: "resetObsTimeZone", GoMethod: "ResetObsTimeZone"},
			_jsii_.MemberMethod{JsiiMethod: "resetObsTimeZoneId", GoMethod: "ResetObsTimeZoneId"},
			_jsii_.MemberMethod{JsiiMethod: "resetObsTransferPath", GoMethod: "ResetObsTransferPath"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LtsTransferV2LogTransferInfoLogTransferDetailOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.ltsTransferV2.LtsTransferV2LogTransferInfoOutputReference",
		reflect.TypeOf((*LtsTransferV2LogTransferInfoOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "logAgencyTransfer", GoGetter: "LogAgencyTransfer"},
			_jsii_.MemberProperty{JsiiProperty: "logAgencyTransferInput", GoGetter: "LogAgencyTransferInput"},
			_jsii_.MemberProperty{JsiiProperty: "logCreatedAt", GoGetter: "LogCreatedAt"},
			_jsii_.MemberProperty{JsiiProperty: "logStorageFormat", GoGetter: "LogStorageFormat"},
			_jsii_.MemberProperty{JsiiProperty: "logStorageFormatInput", GoGetter: "LogStorageFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "logTransferDetail", GoGetter: "LogTransferDetail"},
			_jsii_.MemberProperty{JsiiProperty: "logTransferDetailInput", GoGetter: "LogTransferDetailInput"},
			_jsii_.MemberProperty{JsiiProperty: "logTransferMode", GoGetter: "LogTransferMode"},
			_jsii_.MemberProperty{JsiiProperty: "logTransferModeInput", GoGetter: "LogTransferModeInput"},
			_jsii_.MemberProperty{JsiiProperty: "logTransferStatus", GoGetter: "LogTransferStatus"},
			_jsii_.MemberProperty{JsiiProperty: "logTransferStatusInput", GoGetter: "LogTransferStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "logTransferType", GoGetter: "LogTransferType"},
			_jsii_.MemberProperty{JsiiProperty: "logTransferTypeInput", GoGetter: "LogTransferTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLogAgencyTransfer", GoMethod: "PutLogAgencyTransfer"},
			_jsii_.MemberMethod{JsiiMethod: "putLogTransferDetail", GoMethod: "PutLogTransferDetail"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogAgencyTransfer", GoMethod: "ResetLogAgencyTransfer"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_LtsTransferV2LogTransferInfoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
