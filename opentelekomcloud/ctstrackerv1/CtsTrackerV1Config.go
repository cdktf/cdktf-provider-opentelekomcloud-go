// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ctstrackerv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CtsTrackerV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/cts_tracker_v1#bucket_name CtsTrackerV1#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/cts_tracker_v1#file_prefix_name CtsTrackerV1#file_prefix_name}.
	FilePrefixName *string `field:"optional" json:"filePrefixName" yaml:"filePrefixName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/cts_tracker_v1#id CtsTrackerV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/cts_tracker_v1#is_lts_enabled CtsTrackerV1#is_lts_enabled}.
	IsLtsEnabled interface{} `field:"optional" json:"isLtsEnabled" yaml:"isLtsEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/cts_tracker_v1#timeouts CtsTrackerV1#timeouts}
	Timeouts *CtsTrackerV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

