// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package disdumptaskv2


type DisDumpTaskV2ObsProcessingSchema struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dis_dump_task_v2#timestamp_name DisDumpTaskV2#timestamp_name}.
	TimestampName *string `field:"required" json:"timestampName" yaml:"timestampName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dis_dump_task_v2#timestamp_type DisDumpTaskV2#timestamp_type}.
	TimestampType *string `field:"required" json:"timestampType" yaml:"timestampType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dis_dump_task_v2#timestamp_format DisDumpTaskV2#timestamp_format}.
	TimestampFormat *string `field:"optional" json:"timestampFormat" yaml:"timestampFormat"`
}

