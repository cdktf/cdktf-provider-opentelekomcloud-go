// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package drstaskv3


type DrsTaskV3LimitSpeed struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/drs_task_v3#end_time DrsTaskV3#end_time}.
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/drs_task_v3#speed DrsTaskV3#speed}.
	Speed *string `field:"required" json:"speed" yaml:"speed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/drs_task_v3#start_time DrsTaskV3#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

