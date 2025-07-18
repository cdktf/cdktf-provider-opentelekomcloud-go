// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csbsbackuppolicyv1


type CsbsBackupPolicyV1ScheduledOperation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/csbs_backup_policy_v1#operation_type CsbsBackupPolicyV1#operation_type}.
	OperationType *string `field:"required" json:"operationType" yaml:"operationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/csbs_backup_policy_v1#trigger_pattern CsbsBackupPolicyV1#trigger_pattern}.
	TriggerPattern *string `field:"required" json:"triggerPattern" yaml:"triggerPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/csbs_backup_policy_v1#day_backups CsbsBackupPolicyV1#day_backups}.
	DayBackups *float64 `field:"optional" json:"dayBackups" yaml:"dayBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/csbs_backup_policy_v1#description CsbsBackupPolicyV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/csbs_backup_policy_v1#enabled CsbsBackupPolicyV1#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/csbs_backup_policy_v1#max_backups CsbsBackupPolicyV1#max_backups}.
	MaxBackups *float64 `field:"optional" json:"maxBackups" yaml:"maxBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/csbs_backup_policy_v1#month_backups CsbsBackupPolicyV1#month_backups}.
	MonthBackups *float64 `field:"optional" json:"monthBackups" yaml:"monthBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/csbs_backup_policy_v1#name CsbsBackupPolicyV1#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/csbs_backup_policy_v1#permanent CsbsBackupPolicyV1#permanent}.
	Permanent interface{} `field:"optional" json:"permanent" yaml:"permanent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/csbs_backup_policy_v1#retention_duration_days CsbsBackupPolicyV1#retention_duration_days}.
	RetentionDurationDays *float64 `field:"optional" json:"retentionDurationDays" yaml:"retentionDurationDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/csbs_backup_policy_v1#timezone CsbsBackupPolicyV1#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/csbs_backup_policy_v1#week_backups CsbsBackupPolicyV1#week_backups}.
	WeekBackups *float64 `field:"optional" json:"weekBackups" yaml:"weekBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/csbs_backup_policy_v1#year_backups CsbsBackupPolicyV1#year_backups}.
	YearBackups *float64 `field:"optional" json:"yearBackups" yaml:"yearBackups"`
}

