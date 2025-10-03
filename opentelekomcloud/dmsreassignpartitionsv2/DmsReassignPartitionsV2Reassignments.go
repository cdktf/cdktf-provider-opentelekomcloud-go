// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsreassignpartitionsv2


type DmsReassignPartitionsV2Reassignments struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/dms_reassign_partitions_v2#topic DmsReassignPartitionsV2#topic}.
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// assignments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/dms_reassign_partitions_v2#assignments DmsReassignPartitionsV2#assignments}
	Assignments interface{} `field:"optional" json:"assignments" yaml:"assignments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/dms_reassign_partitions_v2#brokers DmsReassignPartitionsV2#brokers}.
	Brokers *[]*float64 `field:"optional" json:"brokers" yaml:"brokers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/dms_reassign_partitions_v2#replication_factor DmsReassignPartitionsV2#replication_factor}.
	ReplicationFactor *float64 `field:"optional" json:"replicationFactor" yaml:"replicationFactor"`
}

