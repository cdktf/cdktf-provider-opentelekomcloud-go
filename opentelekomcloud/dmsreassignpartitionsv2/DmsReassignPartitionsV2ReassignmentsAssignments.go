// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsreassignpartitionsv2


type DmsReassignPartitionsV2ReassignmentsAssignments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dms_reassign_partitions_v2#partition DmsReassignPartitionsV2#partition}.
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dms_reassign_partitions_v2#partition_brokers DmsReassignPartitionsV2#partition_brokers}.
	PartitionBrokers *[]*float64 `field:"optional" json:"partitionBrokers" yaml:"partitionBrokers"`
}

