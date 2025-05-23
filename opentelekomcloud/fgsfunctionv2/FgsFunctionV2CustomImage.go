// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsfunctionv2


type FgsFunctionV2CustomImage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/fgs_function_v2#url FgsFunctionV2#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/fgs_function_v2#args FgsFunctionV2#args}.
	Args *string `field:"optional" json:"args" yaml:"args"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/fgs_function_v2#command FgsFunctionV2#command}.
	Command *string `field:"optional" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/fgs_function_v2#working_dir FgsFunctionV2#working_dir}.
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
}

