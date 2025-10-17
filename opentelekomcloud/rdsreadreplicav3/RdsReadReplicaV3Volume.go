// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsreadreplicav3


type RdsReadReplicaV3Volume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/rds_read_replica_v3#type RdsReadReplicaV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/rds_read_replica_v3#disk_encryption_id RdsReadReplicaV3#disk_encryption_id}.
	DiskEncryptionId *string `field:"optional" json:"diskEncryptionId" yaml:"diskEncryptionId"`
}

