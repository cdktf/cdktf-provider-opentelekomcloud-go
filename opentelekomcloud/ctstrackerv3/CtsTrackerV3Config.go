// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ctstrackerv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CtsTrackerV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/cts_tracker_v3#status CtsTrackerV3#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/cts_tracker_v3#bucket_name CtsTrackerV3#bucket_name}.
	BucketName *string `field:"optional" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/cts_tracker_v3#compress_type CtsTrackerV3#compress_type}.
	CompressType *string `field:"optional" json:"compressType" yaml:"compressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/cts_tracker_v3#file_prefix_name CtsTrackerV3#file_prefix_name}.
	FilePrefixName *string `field:"optional" json:"filePrefixName" yaml:"filePrefixName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/cts_tracker_v3#is_lts_enabled CtsTrackerV3#is_lts_enabled}.
	IsLtsEnabled interface{} `field:"optional" json:"isLtsEnabled" yaml:"isLtsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/cts_tracker_v3#is_obs_created CtsTrackerV3#is_obs_created}.
	IsObsCreated interface{} `field:"optional" json:"isObsCreated" yaml:"isObsCreated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/cts_tracker_v3#is_sort_by_service CtsTrackerV3#is_sort_by_service}.
	IsSortByService interface{} `field:"optional" json:"isSortByService" yaml:"isSortByService"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/resources/cts_tracker_v3#timeouts CtsTrackerV3#timeouts}
	Timeouts *CtsTrackerV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

