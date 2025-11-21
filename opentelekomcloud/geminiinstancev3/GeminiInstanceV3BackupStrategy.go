// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminiinstancev3


type GeminiInstanceV3BackupStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/gemini_instance_v3#start_time GeminiInstanceV3#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/gemini_instance_v3#keep_days GeminiInstanceV3#keep_days}.
	KeepDays *float64 `field:"optional" json:"keepDays" yaml:"keepDays"`
}

