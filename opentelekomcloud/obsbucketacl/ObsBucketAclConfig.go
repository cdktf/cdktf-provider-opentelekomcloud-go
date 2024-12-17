// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucketacl

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObsBucketAclConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/obs_bucket_acl#bucket ObsBucketAcl#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// account_permission block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/obs_bucket_acl#account_permission ObsBucketAcl#account_permission}
	AccountPermission interface{} `field:"optional" json:"accountPermission" yaml:"accountPermission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/obs_bucket_acl#id ObsBucketAcl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// log_delivery_user_permission block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/obs_bucket_acl#log_delivery_user_permission ObsBucketAcl#log_delivery_user_permission}
	LogDeliveryUserPermission *ObsBucketAclLogDeliveryUserPermission `field:"optional" json:"logDeliveryUserPermission" yaml:"logDeliveryUserPermission"`
	// owner_permission block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/obs_bucket_acl#owner_permission ObsBucketAcl#owner_permission}
	OwnerPermission *ObsBucketAclOwnerPermission `field:"optional" json:"ownerPermission" yaml:"ownerPermission"`
	// public_permission block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/obs_bucket_acl#public_permission ObsBucketAcl#public_permission}
	PublicPermission *ObsBucketAclPublicPermission `field:"optional" json:"publicPermission" yaml:"publicPermission"`
}

