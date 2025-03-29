// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucketobjectacl


type ObsBucketObjectAclPublicPermission struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/obs_bucket_object_acl#access_to_acl ObsBucketObjectAcl#access_to_acl}.
	AccessToAcl *[]*string `field:"optional" json:"accessToAcl" yaml:"accessToAcl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/obs_bucket_object_acl#access_to_object ObsBucketObjectAcl#access_to_object}.
	AccessToObject *[]*string `field:"optional" json:"accessToObject" yaml:"accessToObject"`
}

