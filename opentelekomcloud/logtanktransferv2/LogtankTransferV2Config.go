// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logtanktransferv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogtankTransferV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/logtank_transfer_v2#log_group_id LogtankTransferV2#log_group_id}.
	LogGroupId *string `field:"required" json:"logGroupId" yaml:"logGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/logtank_transfer_v2#log_stream_ids LogtankTransferV2#log_stream_ids}.
	LogStreamIds *[]*string `field:"required" json:"logStreamIds" yaml:"logStreamIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/logtank_transfer_v2#obs_bucket_name LogtankTransferV2#obs_bucket_name}.
	ObsBucketName *string `field:"required" json:"obsBucketName" yaml:"obsBucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/logtank_transfer_v2#period LogtankTransferV2#period}.
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/logtank_transfer_v2#period_unit LogtankTransferV2#period_unit}.
	PeriodUnit *string `field:"required" json:"periodUnit" yaml:"periodUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/logtank_transfer_v2#storage_format LogtankTransferV2#storage_format}.
	StorageFormat *string `field:"required" json:"storageFormat" yaml:"storageFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/logtank_transfer_v2#dir_prefix_name LogtankTransferV2#dir_prefix_name}.
	DirPrefixName *string `field:"optional" json:"dirPrefixName" yaml:"dirPrefixName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/logtank_transfer_v2#id LogtankTransferV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/logtank_transfer_v2#prefix_name LogtankTransferV2#prefix_name}.
	PrefixName *string `field:"optional" json:"prefixName" yaml:"prefixName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/logtank_transfer_v2#switch_on LogtankTransferV2#switch_on}.
	SwitchOn interface{} `field:"optional" json:"switchOn" yaml:"switchOn"`
}

