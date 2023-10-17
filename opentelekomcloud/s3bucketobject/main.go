// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3bucketobject

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-opentelekomcloud.s3BucketObject.S3BucketObject",
		reflect.TypeOf((*S3BucketObject)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acl", GoGetter: "Acl"},
			_jsii_.MemberProperty{JsiiProperty: "aclInput", GoGetter: "AclInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "bucket", GoGetter: "Bucket"},
			_jsii_.MemberProperty{JsiiProperty: "bucketInput", GoGetter: "BucketInput"},
			_jsii_.MemberProperty{JsiiProperty: "cacheControl", GoGetter: "CacheControl"},
			_jsii_.MemberProperty{JsiiProperty: "cacheControlInput", GoGetter: "CacheControlInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "content", GoGetter: "Content"},
			_jsii_.MemberProperty{JsiiProperty: "contentDisposition", GoGetter: "ContentDisposition"},
			_jsii_.MemberProperty{JsiiProperty: "contentDispositionInput", GoGetter: "ContentDispositionInput"},
			_jsii_.MemberProperty{JsiiProperty: "contentEncoding", GoGetter: "ContentEncoding"},
			_jsii_.MemberProperty{JsiiProperty: "contentEncodingInput", GoGetter: "ContentEncodingInput"},
			_jsii_.MemberProperty{JsiiProperty: "contentInput", GoGetter: "ContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "contentLanguage", GoGetter: "ContentLanguage"},
			_jsii_.MemberProperty{JsiiProperty: "contentLanguageInput", GoGetter: "ContentLanguageInput"},
			_jsii_.MemberProperty{JsiiProperty: "contentType", GoGetter: "ContentType"},
			_jsii_.MemberProperty{JsiiProperty: "contentTypeInput", GoGetter: "ContentTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "etagInput", GoGetter: "EtagInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcl", GoMethod: "ResetAcl"},
			_jsii_.MemberMethod{JsiiMethod: "resetCacheControl", GoMethod: "ResetCacheControl"},
			_jsii_.MemberMethod{JsiiMethod: "resetContent", GoMethod: "ResetContent"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentDisposition", GoMethod: "ResetContentDisposition"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentEncoding", GoMethod: "ResetContentEncoding"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentLanguage", GoMethod: "ResetContentLanguage"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentType", GoMethod: "ResetContentType"},
			_jsii_.MemberMethod{JsiiMethod: "resetEtag", GoMethod: "ResetEtag"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetServerSideEncryption", GoMethod: "ResetServerSideEncryption"},
			_jsii_.MemberMethod{JsiiMethod: "resetSource", GoMethod: "ResetSource"},
			_jsii_.MemberMethod{JsiiMethod: "resetSseKmsKeyId", GoMethod: "ResetSseKmsKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebsiteRedirect", GoMethod: "ResetWebsiteRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideEncryption", GoGetter: "ServerSideEncryption"},
			_jsii_.MemberProperty{JsiiProperty: "serverSideEncryptionInput", GoGetter: "ServerSideEncryptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "sourceInput", GoGetter: "SourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "sseKmsKeyId", GoGetter: "SseKmsKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "sseKmsKeyIdInput", GoGetter: "SseKmsKeyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "versionId", GoGetter: "VersionId"},
			_jsii_.MemberProperty{JsiiProperty: "websiteRedirect", GoGetter: "WebsiteRedirect"},
			_jsii_.MemberProperty{JsiiProperty: "websiteRedirectInput", GoGetter: "WebsiteRedirectInput"},
		},
		func() interface{} {
			j := jsiiProxy_S3BucketObject{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-opentelekomcloud.s3BucketObject.S3BucketObjectConfig",
		reflect.TypeOf((*S3BucketObjectConfig)(nil)).Elem(),
	)
}
