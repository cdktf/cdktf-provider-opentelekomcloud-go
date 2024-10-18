// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsfunctionv2


type FgsFunctionV2ReservedInstances struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/fgs_function_v2#count FgsFunctionV2#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/fgs_function_v2#qualifier_name FgsFunctionV2#qualifier_name}.
	QualifierName *string `field:"required" json:"qualifierName" yaml:"qualifierName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/fgs_function_v2#qualifier_type FgsFunctionV2#qualifier_type}.
	QualifierType *string `field:"required" json:"qualifierType" yaml:"qualifierType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/fgs_function_v2#idle_mode FgsFunctionV2#idle_mode}.
	IdleMode interface{} `field:"optional" json:"idleMode" yaml:"idleMode"`
	// tactics_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/fgs_function_v2#tactics_config FgsFunctionV2#tactics_config}
	TacticsConfig *FgsFunctionV2ReservedInstancesTacticsConfig `field:"optional" json:"tacticsConfig" yaml:"tacticsConfig"`
}

