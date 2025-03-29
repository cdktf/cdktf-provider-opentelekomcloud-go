// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodeattachv3


type CceNodeAttachV3Taints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/cce_node_attach_v3#effect CceNodeAttachV3#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/cce_node_attach_v3#key CceNodeAttachV3#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.34/docs/resources/cce_node_attach_v3#value CceNodeAttachV3#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

