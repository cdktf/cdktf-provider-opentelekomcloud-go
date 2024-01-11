// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cbrpolicyv3


type CbrPolicyV3OperationDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/cbr_policy_v3#timezone CbrPolicyV3#timezone}.
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/cbr_policy_v3#day_backups CbrPolicyV3#day_backups}.
	DayBackups *float64 `field:"optional" json:"dayBackups" yaml:"dayBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/cbr_policy_v3#max_backups CbrPolicyV3#max_backups}.
	MaxBackups *float64 `field:"optional" json:"maxBackups" yaml:"maxBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/cbr_policy_v3#month_backups CbrPolicyV3#month_backups}.
	MonthBackups *float64 `field:"optional" json:"monthBackups" yaml:"monthBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/cbr_policy_v3#retention_duration_days CbrPolicyV3#retention_duration_days}.
	RetentionDurationDays *float64 `field:"optional" json:"retentionDurationDays" yaml:"retentionDurationDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/cbr_policy_v3#week_backups CbrPolicyV3#week_backups}.
	WeekBackups *float64 `field:"optional" json:"weekBackups" yaml:"weekBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/cbr_policy_v3#year_backups CbrPolicyV3#year_backups}.
	YearBackups *float64 `field:"optional" json:"yearBackups" yaml:"yearBackups"`
}

