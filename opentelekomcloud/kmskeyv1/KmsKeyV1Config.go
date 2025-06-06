// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmskeyv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KmsKeyV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/kms_key_v1#key_alias KmsKeyV1#key_alias}.
	KeyAlias *string `field:"required" json:"keyAlias" yaml:"keyAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/kms_key_v1#allow_cancel_deletion KmsKeyV1#allow_cancel_deletion}.
	AllowCancelDeletion interface{} `field:"optional" json:"allowCancelDeletion" yaml:"allowCancelDeletion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/kms_key_v1#id KmsKeyV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/kms_key_v1#is_enabled KmsKeyV1#is_enabled}.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/kms_key_v1#key_description KmsKeyV1#key_description}.
	KeyDescription *string `field:"optional" json:"keyDescription" yaml:"keyDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/kms_key_v1#origin KmsKeyV1#origin}.
	Origin *string `field:"optional" json:"origin" yaml:"origin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/kms_key_v1#pending_days KmsKeyV1#pending_days}.
	PendingDays *string `field:"optional" json:"pendingDays" yaml:"pendingDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/kms_key_v1#realm KmsKeyV1#realm}.
	Realm *string `field:"optional" json:"realm" yaml:"realm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/kms_key_v1#rotation_enabled KmsKeyV1#rotation_enabled}.
	RotationEnabled interface{} `field:"optional" json:"rotationEnabled" yaml:"rotationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/kms_key_v1#rotation_interval KmsKeyV1#rotation_interval}.
	RotationInterval *float64 `field:"optional" json:"rotationInterval" yaml:"rotationInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/kms_key_v1#tags KmsKeyV1#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

