// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodeattachv3


type CceNodeAttachV3Storage struct {
	// groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cce_node_attach_v3#groups CceNodeAttachV3#groups}
	Groups interface{} `field:"required" json:"groups" yaml:"groups"`
	// selectors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/cce_node_attach_v3#selectors CceNodeAttachV3#selectors}
	Selectors interface{} `field:"required" json:"selectors" yaml:"selectors"`
}

