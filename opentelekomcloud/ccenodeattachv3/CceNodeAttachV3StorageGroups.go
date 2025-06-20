// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodeattachv3


type CceNodeAttachV3StorageGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/cce_node_attach_v3#name CceNodeAttachV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/cce_node_attach_v3#selector_names CceNodeAttachV3#selector_names}.
	SelectorNames *[]*string `field:"required" json:"selectorNames" yaml:"selectorNames"`
	// virtual_spaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/cce_node_attach_v3#virtual_spaces CceNodeAttachV3#virtual_spaces}
	VirtualSpaces interface{} `field:"required" json:"virtualSpaces" yaml:"virtualSpaces"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/cce_node_attach_v3#cce_managed CceNodeAttachV3#cce_managed}.
	CceManaged interface{} `field:"optional" json:"cceManaged" yaml:"cceManaged"`
}

