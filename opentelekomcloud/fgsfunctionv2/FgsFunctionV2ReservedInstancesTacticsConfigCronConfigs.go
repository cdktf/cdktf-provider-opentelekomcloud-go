// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsfunctionv2


type FgsFunctionV2ReservedInstancesTacticsConfigCronConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/fgs_function_v2#count FgsFunctionV2#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/fgs_function_v2#cron FgsFunctionV2#cron}.
	Cron *string `field:"required" json:"cron" yaml:"cron"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/fgs_function_v2#expired_time FgsFunctionV2#expired_time}.
	ExpiredTime *float64 `field:"required" json:"expiredTime" yaml:"expiredTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/fgs_function_v2#name FgsFunctionV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/fgs_function_v2#start_time FgsFunctionV2#start_time}.
	StartTime *float64 `field:"required" json:"startTime" yaml:"startTime"`
}

