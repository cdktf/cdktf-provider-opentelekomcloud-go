// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudhssintrusioneventsv5

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "beginTime", GoGetter: "BeginTime"},
			_jsii_.MemberProperty{JsiiProperty: "beginTimeInput", GoGetter: "BeginTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "category", GoGetter: "Category"},
			_jsii_.MemberProperty{JsiiProperty: "categoryInput", GoGetter: "CategoryInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "containerName", GoGetter: "ContainerName"},
			_jsii_.MemberProperty{JsiiProperty: "containerNameInput", GoGetter: "ContainerNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "days", GoGetter: "Days"},
			_jsii_.MemberProperty{JsiiProperty: "daysInput", GoGetter: "DaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "endTime", GoGetter: "EndTime"},
			_jsii_.MemberProperty{JsiiProperty: "endTimeInput", GoGetter: "EndTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "enterpriseProjectId", GoGetter: "EnterpriseProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "enterpriseProjectIdInput", GoGetter: "EnterpriseProjectIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "events", GoGetter: "Events"},
			_jsii_.MemberProperty{JsiiProperty: "eventTypes", GoGetter: "EventTypes"},
			_jsii_.MemberProperty{JsiiProperty: "eventTypesInput", GoGetter: "EventTypesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "handleStatus", GoGetter: "HandleStatus"},
			_jsii_.MemberProperty{JsiiProperty: "handleStatusInput", GoGetter: "HandleStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "hostId", GoGetter: "HostId"},
			_jsii_.MemberProperty{JsiiProperty: "hostIdInput", GoGetter: "HostIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "hostName", GoGetter: "HostName"},
			_jsii_.MemberProperty{JsiiProperty: "hostNameInput", GoGetter: "HostNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "privateIp", GoGetter: "PrivateIp"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpInput", GoGetter: "PrivateIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberMethod{JsiiMethod: "resetBeginTime", GoMethod: "ResetBeginTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetContainerName", GoMethod: "ResetContainerName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDays", GoMethod: "ResetDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndTime", GoMethod: "ResetEndTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnterpriseProjectId", GoMethod: "ResetEnterpriseProjectId"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventTypes", GoMethod: "ResetEventTypes"},
			_jsii_.MemberMethod{JsiiMethod: "resetHandleStatus", GoMethod: "ResetHandleStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostId", GoMethod: "ResetHostId"},
			_jsii_.MemberMethod{JsiiMethod: "resetHostName", GoMethod: "ResetHostName"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateIp", GoMethod: "ResetPrivateIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetSeverity", GoMethod: "ResetSeverity"},
			_jsii_.MemberProperty{JsiiProperty: "severity", GoGetter: "Severity"},
			_jsii_.MemberProperty{JsiiProperty: "severityInput", GoGetter: "SeverityInput"},
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
			j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5Config",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5Config)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5Events",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5Events)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStruct",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStruct)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructList",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fdCount", GoGetter: "FdCount"},
			_jsii_.MemberProperty{JsiiProperty: "fdInfo", GoGetter: "FdInfo"},
			_jsii_.MemberProperty{JsiiProperty: "fileAction", GoGetter: "FileAction"},
			_jsii_.MemberProperty{JsiiProperty: "fileAlias", GoGetter: "FileAlias"},
			_jsii_.MemberProperty{JsiiProperty: "fileAtime", GoGetter: "FileAtime"},
			_jsii_.MemberProperty{JsiiProperty: "fileAttr", GoGetter: "FileAttr"},
			_jsii_.MemberProperty{JsiiProperty: "fileChangeAttr", GoGetter: "FileChangeAttr"},
			_jsii_.MemberProperty{JsiiProperty: "fileContent", GoGetter: "FileContent"},
			_jsii_.MemberProperty{JsiiProperty: "fileCtime", GoGetter: "FileCtime"},
			_jsii_.MemberProperty{JsiiProperty: "fileDesc", GoGetter: "FileDesc"},
			_jsii_.MemberProperty{JsiiProperty: "fileHash", GoGetter: "FileHash"},
			_jsii_.MemberProperty{JsiiProperty: "fileKeyWord", GoGetter: "FileKeyWord"},
			_jsii_.MemberProperty{JsiiProperty: "fileMd5", GoGetter: "FileMd5"},
			_jsii_.MemberProperty{JsiiProperty: "fileMtime", GoGetter: "FileMtime"},
			_jsii_.MemberProperty{JsiiProperty: "fileNewPath", GoGetter: "FileNewPath"},
			_jsii_.MemberProperty{JsiiProperty: "fileOperation", GoGetter: "FileOperation"},
			_jsii_.MemberProperty{JsiiProperty: "filePath", GoGetter: "FilePath"},
			_jsii_.MemberProperty{JsiiProperty: "fileSha256", GoGetter: "FileSha256"},
			_jsii_.MemberProperty{JsiiProperty: "fileSize", GoGetter: "FileSize"},
			_jsii_.MemberProperty{JsiiProperty: "fileType", GoGetter: "FileType"},
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
			_jsii_.MemberProperty{JsiiProperty: "isDir", GoGetter: "IsDir"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsFileInfoListStructOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsList",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsOperateDetailListStruct",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsOperateDetailListStruct)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsOperateDetailListStructList",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsOperateDetailListStructList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOperateDetailListStructList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsOperateDetailListStructOutputReference",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsOperateDetailListStructOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentId", GoGetter: "AgentId"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileAttr", GoGetter: "FileAttr"},
			_jsii_.MemberProperty{JsiiProperty: "fileHash", GoGetter: "FileHash"},
			_jsii_.MemberProperty{JsiiProperty: "filePath", GoGetter: "FilePath"},
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
			_jsii_.MemberProperty{JsiiProperty: "hash", GoGetter: "Hash"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isParent", GoGetter: "IsParent"},
			_jsii_.MemberProperty{JsiiProperty: "keyword", GoGetter: "Keyword"},
			_jsii_.MemberProperty{JsiiProperty: "loginIp", GoGetter: "LoginIp"},
			_jsii_.MemberProperty{JsiiProperty: "loginUserName", GoGetter: "LoginUserName"},
			_jsii_.MemberProperty{JsiiProperty: "privateIp", GoGetter: "PrivateIp"},
			_jsii_.MemberProperty{JsiiProperty: "processPid", GoGetter: "ProcessPid"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOperateDetailListStructOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "agentStatus", GoGetter: "AgentStatus"},
			_jsii_.MemberProperty{JsiiProperty: "assetValue", GoGetter: "AssetValue"},
			_jsii_.MemberProperty{JsiiProperty: "attackPhase", GoGetter: "AttackPhase"},
			_jsii_.MemberProperty{JsiiProperty: "attackTag", GoGetter: "AttackTag"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "containerName", GoGetter: "ContainerName"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "eventClassId", GoGetter: "EventClassId"},
			_jsii_.MemberProperty{JsiiProperty: "eventDetails", GoGetter: "EventDetails"},
			_jsii_.MemberProperty{JsiiProperty: "eventName", GoGetter: "EventName"},
			_jsii_.MemberProperty{JsiiProperty: "eventType", GoGetter: "EventType"},
			_jsii_.MemberProperty{JsiiProperty: "fileInfoList", GoGetter: "FileInfoList"},
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
			_jsii_.MemberProperty{JsiiProperty: "handleMethod", GoGetter: "HandleMethod"},
			_jsii_.MemberProperty{JsiiProperty: "handler", GoGetter: "Handler"},
			_jsii_.MemberProperty{JsiiProperty: "handleStatus", GoGetter: "HandleStatus"},
			_jsii_.MemberProperty{JsiiProperty: "handleTime", GoGetter: "HandleTime"},
			_jsii_.MemberProperty{JsiiProperty: "hostId", GoGetter: "HostId"},
			_jsii_.MemberProperty{JsiiProperty: "hostName", GoGetter: "HostName"},
			_jsii_.MemberProperty{JsiiProperty: "hostStatus", GoGetter: "HostStatus"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "imageName", GoGetter: "ImageName"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "occurTime", GoGetter: "OccurTime"},
			_jsii_.MemberProperty{JsiiProperty: "operateAcceptList", GoGetter: "OperateAcceptList"},
			_jsii_.MemberProperty{JsiiProperty: "operateDetailList", GoGetter: "OperateDetailList"},
			_jsii_.MemberProperty{JsiiProperty: "osType", GoGetter: "OsType"},
			_jsii_.MemberProperty{JsiiProperty: "privateIp", GoGetter: "PrivateIp"},
			_jsii_.MemberProperty{JsiiProperty: "processInfoList", GoGetter: "ProcessInfoList"},
			_jsii_.MemberProperty{JsiiProperty: "protectStatus", GoGetter: "ProtectStatus"},
			_jsii_.MemberProperty{JsiiProperty: "publicIp", GoGetter: "PublicIp"},
			_jsii_.MemberProperty{JsiiProperty: "recommendation", GoGetter: "Recommendation"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resourceInfo", GoGetter: "ResourceInfo"},
			_jsii_.MemberProperty{JsiiProperty: "severity", GoGetter: "Severity"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userInfoList", GoGetter: "UserInfoList"},
		},
		func() interface{} {
			j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStruct",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStruct)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructList",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "childProcessCmdline", GoGetter: "ChildProcessCmdline"},
			_jsii_.MemberProperty{JsiiProperty: "childProcessEgid", GoGetter: "ChildProcessEgid"},
			_jsii_.MemberProperty{JsiiProperty: "childProcessEuid", GoGetter: "ChildProcessEuid"},
			_jsii_.MemberProperty{JsiiProperty: "childProcessFilename", GoGetter: "ChildProcessFilename"},
			_jsii_.MemberProperty{JsiiProperty: "childProcessGid", GoGetter: "ChildProcessGid"},
			_jsii_.MemberProperty{JsiiProperty: "childProcessName", GoGetter: "ChildProcessName"},
			_jsii_.MemberProperty{JsiiProperty: "childProcessPath", GoGetter: "ChildProcessPath"},
			_jsii_.MemberProperty{JsiiProperty: "childProcessPid", GoGetter: "ChildProcessPid"},
			_jsii_.MemberProperty{JsiiProperty: "childProcessStartTime", GoGetter: "ChildProcessStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "childProcessUid", GoGetter: "ChildProcessUid"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "escapeCmd", GoGetter: "EscapeCmd"},
			_jsii_.MemberProperty{JsiiProperty: "escapeMode", GoGetter: "EscapeMode"},
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
			_jsii_.MemberProperty{JsiiProperty: "parentProcessCmdline", GoGetter: "ParentProcessCmdline"},
			_jsii_.MemberProperty{JsiiProperty: "parentProcessEgid", GoGetter: "ParentProcessEgid"},
			_jsii_.MemberProperty{JsiiProperty: "parentProcessEuid", GoGetter: "ParentProcessEuid"},
			_jsii_.MemberProperty{JsiiProperty: "parentProcessFilename", GoGetter: "ParentProcessFilename"},
			_jsii_.MemberProperty{JsiiProperty: "parentProcessGid", GoGetter: "ParentProcessGid"},
			_jsii_.MemberProperty{JsiiProperty: "parentProcessName", GoGetter: "ParentProcessName"},
			_jsii_.MemberProperty{JsiiProperty: "parentProcessPath", GoGetter: "ParentProcessPath"},
			_jsii_.MemberProperty{JsiiProperty: "parentProcessPid", GoGetter: "ParentProcessPid"},
			_jsii_.MemberProperty{JsiiProperty: "parentProcessStartTime", GoGetter: "ParentProcessStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "parentProcessUid", GoGetter: "ParentProcessUid"},
			_jsii_.MemberProperty{JsiiProperty: "processCmdline", GoGetter: "ProcessCmdline"},
			_jsii_.MemberProperty{JsiiProperty: "processEgid", GoGetter: "ProcessEgid"},
			_jsii_.MemberProperty{JsiiProperty: "processEuid", GoGetter: "ProcessEuid"},
			_jsii_.MemberProperty{JsiiProperty: "processFilename", GoGetter: "ProcessFilename"},
			_jsii_.MemberProperty{JsiiProperty: "processGid", GoGetter: "ProcessGid"},
			_jsii_.MemberProperty{JsiiProperty: "processHash", GoGetter: "ProcessHash"},
			_jsii_.MemberProperty{JsiiProperty: "processName", GoGetter: "ProcessName"},
			_jsii_.MemberProperty{JsiiProperty: "processPath", GoGetter: "ProcessPath"},
			_jsii_.MemberProperty{JsiiProperty: "processPid", GoGetter: "ProcessPid"},
			_jsii_.MemberProperty{JsiiProperty: "processStartTime", GoGetter: "ProcessStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "processUid", GoGetter: "ProcessUid"},
			_jsii_.MemberProperty{JsiiProperty: "processUsername", GoGetter: "ProcessUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "virtCmd", GoGetter: "VirtCmd"},
			_jsii_.MemberProperty{JsiiProperty: "virtProcessName", GoGetter: "VirtProcessName"},
		},
		func() interface{} {
			j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsProcessInfoListStructOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsResourceInfo",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsResourceInfo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsResourceInfoList",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsResourceInfoList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsResourceInfoList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsResourceInfoOutputReference",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsResourceInfoOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "containerId", GoGetter: "ContainerId"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "domainId", GoGetter: "DomainId"},
			_jsii_.MemberProperty{JsiiProperty: "ecsId", GoGetter: "EcsId"},
			_jsii_.MemberProperty{JsiiProperty: "enterpriseProjectId", GoGetter: "EnterpriseProjectId"},
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
			_jsii_.MemberProperty{JsiiProperty: "hostAttr", GoGetter: "HostAttr"},
			_jsii_.MemberProperty{JsiiProperty: "imageId", GoGetter: "ImageId"},
			_jsii_.MemberProperty{JsiiProperty: "imageName", GoGetter: "ImageName"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "microservice", GoGetter: "Microservice"},
			_jsii_.MemberProperty{JsiiProperty: "osBit", GoGetter: "OsBit"},
			_jsii_.MemberProperty{JsiiProperty: "osName", GoGetter: "OsName"},
			_jsii_.MemberProperty{JsiiProperty: "osType", GoGetter: "OsType"},
			_jsii_.MemberProperty{JsiiProperty: "osVersion", GoGetter: "OsVersion"},
			_jsii_.MemberProperty{JsiiProperty: "projectId", GoGetter: "ProjectId"},
			_jsii_.MemberProperty{JsiiProperty: "regionName", GoGetter: "RegionName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "service", GoGetter: "Service"},
			_jsii_.MemberProperty{JsiiProperty: "sysArch", GoGetter: "SysArch"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vmName", GoGetter: "VmName"},
			_jsii_.MemberProperty{JsiiProperty: "vmUuid", GoGetter: "VmUuid"},
			_jsii_.MemberProperty{JsiiProperty: "vpcId", GoGetter: "VpcId"},
		},
		func() interface{} {
			j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsResourceInfoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsUserInfoListStruct",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsUserInfoListStruct)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsUserInfoListStructList",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsUserInfoListStructList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
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
			j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsUserInfoListStructList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.dataOpentelekomcloudHssIntrusionEventsV5.DataOpentelekomcloudHssIntrusionEventsV5EventsUserInfoListStructOutputReference",
		reflect.TypeOf((*DataOpentelekomcloudHssIntrusionEventsV5EventsUserInfoListStructOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "loginFailCount", GoGetter: "LoginFailCount"},
			_jsii_.MemberProperty{JsiiProperty: "loginIp", GoGetter: "LoginIp"},
			_jsii_.MemberProperty{JsiiProperty: "loginLastTime", GoGetter: "LoginLastTime"},
			_jsii_.MemberProperty{JsiiProperty: "loginMode", GoGetter: "LoginMode"},
			_jsii_.MemberProperty{JsiiProperty: "pwdHash", GoGetter: "PwdHash"},
			_jsii_.MemberProperty{JsiiProperty: "pwdMaxDays", GoGetter: "PwdMaxDays"},
			_jsii_.MemberProperty{JsiiProperty: "pwdMinDays", GoGetter: "PwdMinDays"},
			_jsii_.MemberProperty{JsiiProperty: "pwdUsedDays", GoGetter: "PwdUsedDays"},
			_jsii_.MemberProperty{JsiiProperty: "pwdWarnLeftDays", GoGetter: "PwdWarnLeftDays"},
			_jsii_.MemberProperty{JsiiProperty: "pwdWithFuzzing", GoGetter: "PwdWithFuzzing"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "servicePort", GoGetter: "ServicePort"},
			_jsii_.MemberProperty{JsiiProperty: "serviceType", GoGetter: "ServiceType"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "userGid", GoGetter: "UserGid"},
			_jsii_.MemberProperty{JsiiProperty: "userGroupName", GoGetter: "UserGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "userHomeDir", GoGetter: "UserHomeDir"},
			_jsii_.MemberProperty{JsiiProperty: "userId", GoGetter: "UserId"},
			_jsii_.MemberProperty{JsiiProperty: "userName", GoGetter: "UserName"},
		},
		func() interface{} {
			j := jsiiProxy_DataOpentelekomcloudHssIntrusionEventsV5EventsUserInfoListStructOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
