// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsinstancev3


type RdsInstanceV3RestorePoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/rds_instance_v3#instance_id RdsInstanceV3#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/rds_instance_v3#backup_id RdsInstanceV3#backup_id}.
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/rds_instance_v3#restore_time RdsInstanceV3#restore_time}.
	RestoreTime *float64 `field:"optional" json:"restoreTime" yaml:"restoreTime"`
}

