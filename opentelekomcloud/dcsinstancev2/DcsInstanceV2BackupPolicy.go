// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcsinstancev2


type DcsInstanceV2BackupPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/dcs_instance_v2#backup_at DcsInstanceV2#backup_at}.
	BackupAt *[]*float64 `field:"required" json:"backupAt" yaml:"backupAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/dcs_instance_v2#begin_at DcsInstanceV2#begin_at}.
	BeginAt *string `field:"required" json:"beginAt" yaml:"beginAt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/dcs_instance_v2#backup_type DcsInstanceV2#backup_type}.
	BackupType *string `field:"optional" json:"backupType" yaml:"backupType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/dcs_instance_v2#period_type DcsInstanceV2#period_type}.
	PeriodType *string `field:"optional" json:"periodType" yaml:"periodType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/dcs_instance_v2#save_days DcsInstanceV2#save_days}.
	SaveDays *float64 `field:"optional" json:"saveDays" yaml:"saveDays"`
}

