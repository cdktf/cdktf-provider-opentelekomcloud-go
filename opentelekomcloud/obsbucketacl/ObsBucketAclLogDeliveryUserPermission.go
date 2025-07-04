// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucketacl


type ObsBucketAclLogDeliveryUserPermission struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/obs_bucket_acl#access_to_acl ObsBucketAcl#access_to_acl}.
	AccessToAcl *[]*string `field:"optional" json:"accessToAcl" yaml:"accessToAcl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/obs_bucket_acl#access_to_bucket ObsBucketAcl#access_to_bucket}.
	AccessToBucket *[]*string `field:"optional" json:"accessToBucket" yaml:"accessToBucket"`
}

