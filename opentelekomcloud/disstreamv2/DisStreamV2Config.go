// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package disstreamv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DisStreamV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/dis_stream_v2#name DisStreamV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/dis_stream_v2#partition_count DisStreamV2#partition_count}.
	PartitionCount *float64 `field:"required" json:"partitionCount" yaml:"partitionCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/dis_stream_v2#auto_scale_max_partition_count DisStreamV2#auto_scale_max_partition_count}.
	AutoScaleMaxPartitionCount *float64 `field:"optional" json:"autoScaleMaxPartitionCount" yaml:"autoScaleMaxPartitionCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/dis_stream_v2#auto_scale_min_partition_count DisStreamV2#auto_scale_min_partition_count}.
	AutoScaleMinPartitionCount *float64 `field:"optional" json:"autoScaleMinPartitionCount" yaml:"autoScaleMinPartitionCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/dis_stream_v2#compression_format DisStreamV2#compression_format}.
	CompressionFormat *string `field:"optional" json:"compressionFormat" yaml:"compressionFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/dis_stream_v2#data_type DisStreamV2#data_type}.
	DataType *string `field:"optional" json:"dataType" yaml:"dataType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/dis_stream_v2#id DisStreamV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/dis_stream_v2#retention_period DisStreamV2#retention_period}.
	RetentionPeriod *float64 `field:"optional" json:"retentionPeriod" yaml:"retentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/dis_stream_v2#stream_type DisStreamV2#stream_type}.
	StreamType *string `field:"optional" json:"streamType" yaml:"streamType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/dis_stream_v2#tags DisStreamV2#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.5/docs/resources/dis_stream_v2#timeouts DisStreamV2#timeouts}
	Timeouts *DisStreamV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

