// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsinstancev3


type RdsInstanceV3Volume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.16/docs/resources/rds_instance_v3#size RdsInstanceV3#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.16/docs/resources/rds_instance_v3#type RdsInstanceV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.16/docs/resources/rds_instance_v3#disk_encryption_id RdsInstanceV3#disk_encryption_id}.
	DiskEncryptionId *string `field:"optional" json:"diskEncryptionId" yaml:"diskEncryptionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.16/docs/resources/rds_instance_v3#limit_size RdsInstanceV3#limit_size}.
	LimitSize *float64 `field:"optional" json:"limitSize" yaml:"limitSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.16/docs/resources/rds_instance_v3#trigger_threshold RdsInstanceV3#trigger_threshold}.
	TriggerThreshold *float64 `field:"optional" json:"triggerThreshold" yaml:"triggerThreshold"`
}

