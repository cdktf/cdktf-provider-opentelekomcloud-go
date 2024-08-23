// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package erpropagationv3


type ErPropagationV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.17/docs/resources/er_propagation_v3#create ErPropagationV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.17/docs/resources/er_propagation_v3#delete ErPropagationV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

