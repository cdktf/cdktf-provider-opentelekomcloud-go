// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodev3


type CceNodeV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/cce_node_v3#create CceNodeV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/cce_node_v3#delete CceNodeV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

