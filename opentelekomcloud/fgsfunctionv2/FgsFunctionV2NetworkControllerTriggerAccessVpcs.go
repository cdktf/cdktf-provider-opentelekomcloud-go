// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fgsfunctionv2


type FgsFunctionV2NetworkControllerTriggerAccessVpcs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/fgs_function_v2#vpc_id FgsFunctionV2#vpc_id}.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/fgs_function_v2#vpc_name FgsFunctionV2#vpc_name}.
	VpcName *string `field:"optional" json:"vpcName" yaml:"vpcName"`
}

