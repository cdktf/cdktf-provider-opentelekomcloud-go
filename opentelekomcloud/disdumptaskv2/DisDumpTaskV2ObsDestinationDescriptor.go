// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package disdumptaskv2


type DisDumpTaskV2ObsDestinationDescriptor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/dis_dump_task_v2#agency_name DisDumpTaskV2#agency_name}.
	AgencyName *string `field:"required" json:"agencyName" yaml:"agencyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/dis_dump_task_v2#deliver_time_interval DisDumpTaskV2#deliver_time_interval}.
	DeliverTimeInterval *float64 `field:"required" json:"deliverTimeInterval" yaml:"deliverTimeInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/dis_dump_task_v2#obs_bucket_path DisDumpTaskV2#obs_bucket_path}.
	ObsBucketPath *string `field:"required" json:"obsBucketPath" yaml:"obsBucketPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/dis_dump_task_v2#task_name DisDumpTaskV2#task_name}.
	TaskName *string `field:"required" json:"taskName" yaml:"taskName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/dis_dump_task_v2#consumer_strategy DisDumpTaskV2#consumer_strategy}.
	ConsumerStrategy *string `field:"optional" json:"consumerStrategy" yaml:"consumerStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/dis_dump_task_v2#destination_file_type DisDumpTaskV2#destination_file_type}.
	DestinationFileType *string `field:"optional" json:"destinationFileType" yaml:"destinationFileType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/dis_dump_task_v2#file_prefix DisDumpTaskV2#file_prefix}.
	FilePrefix *string `field:"optional" json:"filePrefix" yaml:"filePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/dis_dump_task_v2#partition_format DisDumpTaskV2#partition_format}.
	PartitionFormat *string `field:"optional" json:"partitionFormat" yaml:"partitionFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.10/docs/resources/dis_dump_task_v2#record_delimiter DisDumpTaskV2#record_delimiter}.
	RecordDelimiter *string `field:"optional" json:"recordDelimiter" yaml:"recordDelimiter"`
}

