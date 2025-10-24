// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsfunctionv2


type FgsFunctionV2Versions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/fgs_function_v2#name FgsFunctionV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// aliases block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/fgs_function_v2#aliases FgsFunctionV2#aliases}
	Aliases *FgsFunctionV2VersionsAliases `field:"optional" json:"aliases" yaml:"aliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/fgs_function_v2#description FgsFunctionV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

