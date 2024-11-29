// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmskeyv1

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.kmsKeyV1.KmsKeyV1",
		reflect.TypeOf((*KmsKeyV1)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowCancelDeletion", GoGetter: "AllowCancelDeletion"},
			_jsii_.MemberProperty{JsiiProperty: "allowCancelDeletionInput", GoGetter: "AllowCancelDeletionInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "creationDate", GoGetter: "CreationDate"},
			_jsii_.MemberProperty{JsiiProperty: "defaultKeyFlag", GoGetter: "DefaultKeyFlag"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "domainId", GoGetter: "DomainId"},
			_jsii_.MemberProperty{JsiiProperty: "expirationTime", GoGetter: "ExpirationTime"},
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
			_jsii_.MemberProperty{JsiiProperty: "isEnabled", GoGetter: "IsEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "isEnabledInput", GoGetter: "IsEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyAlias", GoGetter: "KeyAlias"},
			_jsii_.MemberProperty{JsiiProperty: "keyAliasInput", GoGetter: "KeyAliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyDescription", GoGetter: "KeyDescription"},
			_jsii_.MemberProperty{JsiiProperty: "keyDescriptionInput", GoGetter: "KeyDescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyState", GoGetter: "KeyState"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "origin", GoGetter: "Origin"},
			_jsii_.MemberProperty{JsiiProperty: "originInput", GoGetter: "OriginInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pendingDays", GoGetter: "PendingDays"},
			_jsii_.MemberProperty{JsiiProperty: "pendingDaysInput", GoGetter: "PendingDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "realm", GoGetter: "Realm"},
			_jsii_.MemberProperty{JsiiProperty: "realmInput", GoGetter: "RealmInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowCancelDeletion", GoMethod: "ResetAllowCancelDeletion"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsEnabled", GoMethod: "ResetIsEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyDescription", GoMethod: "ResetKeyDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrigin", GoMethod: "ResetOrigin"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPendingDays", GoMethod: "ResetPendingDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetRealm", GoMethod: "ResetRealm"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationEnabled", GoMethod: "ResetRotationEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetRotationInterval", GoMethod: "ResetRotationInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberProperty{JsiiProperty: "rotationEnabled", GoGetter: "RotationEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "rotationEnabledInput", GoGetter: "RotationEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "rotationInterval", GoGetter: "RotationInterval"},
			_jsii_.MemberProperty{JsiiProperty: "rotationIntervalInput", GoGetter: "RotationIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "rotationNumber", GoGetter: "RotationNumber"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledDeletionDate", GoGetter: "ScheduledDeletionDate"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_KmsKeyV1{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.kmsKeyV1.KmsKeyV1Config",
		reflect.TypeOf((*KmsKeyV1Config)(nil)).Elem(),
	)
}
