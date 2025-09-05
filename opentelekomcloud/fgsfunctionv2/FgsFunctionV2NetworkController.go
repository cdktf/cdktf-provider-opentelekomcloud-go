// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsfunctionv2


type FgsFunctionV2NetworkController struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/fgs_function_v2#disable_public_network FgsFunctionV2#disable_public_network}.
	DisablePublicNetwork interface{} `field:"optional" json:"disablePublicNetwork" yaml:"disablePublicNetwork"`
	// trigger_access_vpcs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/fgs_function_v2#trigger_access_vpcs FgsFunctionV2#trigger_access_vpcs}
	TriggerAccessVpcs interface{} `field:"optional" json:"triggerAccessVpcs" yaml:"triggerAccessVpcs"`
}

