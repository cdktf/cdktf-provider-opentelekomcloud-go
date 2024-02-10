// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package discheckpointv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DisCheckpointV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/dis_checkpoint_v2#app_name DisCheckpointV2#app_name}.
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/dis_checkpoint_v2#partition_id DisCheckpointV2#partition_id}.
	PartitionId *string `field:"required" json:"partitionId" yaml:"partitionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/dis_checkpoint_v2#sequence_number DisCheckpointV2#sequence_number}.
	SequenceNumber *string `field:"required" json:"sequenceNumber" yaml:"sequenceNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/dis_checkpoint_v2#stream_name DisCheckpointV2#stream_name}.
	StreamName *string `field:"required" json:"streamName" yaml:"streamName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/dis_checkpoint_v2#checkpoint_type DisCheckpointV2#checkpoint_type}.
	CheckpointType *string `field:"optional" json:"checkpointType" yaml:"checkpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/dis_checkpoint_v2#id DisCheckpointV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/dis_checkpoint_v2#metadata DisCheckpointV2#metadata}.
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/dis_checkpoint_v2#timeouts DisCheckpointV2#timeouts}
	Timeouts *DisCheckpointV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

