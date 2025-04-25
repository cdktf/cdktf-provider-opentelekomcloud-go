// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ddsinstancev3


type DdsInstanceV3BackupStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dds_instance_v3#keep_days DdsInstanceV3#keep_days}.
	KeepDays *float64 `field:"required" json:"keepDays" yaml:"keepDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dds_instance_v3#start_time DdsInstanceV3#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dds_instance_v3#period DdsInstanceV3#period}.
	Period *string `field:"optional" json:"period" yaml:"period"`
}

