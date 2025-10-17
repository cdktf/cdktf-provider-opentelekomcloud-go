// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudkmskeyv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudKmsKeyV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/kms_key_v1#default_key_flag DataOpentelekomcloudKmsKeyV1#default_key_flag}.
	DefaultKeyFlag *string `field:"optional" json:"defaultKeyFlag" yaml:"defaultKeyFlag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/kms_key_v1#domain_id DataOpentelekomcloudKmsKeyV1#domain_id}.
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/kms_key_v1#id DataOpentelekomcloudKmsKeyV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/kms_key_v1#key_alias DataOpentelekomcloudKmsKeyV1#key_alias}.
	KeyAlias *string `field:"optional" json:"keyAlias" yaml:"keyAlias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/kms_key_v1#key_description DataOpentelekomcloudKmsKeyV1#key_description}.
	KeyDescription *string `field:"optional" json:"keyDescription" yaml:"keyDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/kms_key_v1#key_id DataOpentelekomcloudKmsKeyV1#key_id}.
	KeyId *string `field:"optional" json:"keyId" yaml:"keyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/kms_key_v1#key_state DataOpentelekomcloudKmsKeyV1#key_state}.
	KeyState *string `field:"optional" json:"keyState" yaml:"keyState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/kms_key_v1#origin DataOpentelekomcloudKmsKeyV1#origin}.
	Origin *string `field:"optional" json:"origin" yaml:"origin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/kms_key_v1#realm DataOpentelekomcloudKmsKeyV1#realm}.
	Realm *string `field:"optional" json:"realm" yaml:"realm"`
}

