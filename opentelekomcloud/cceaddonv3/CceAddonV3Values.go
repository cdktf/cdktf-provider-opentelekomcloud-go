// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cceaddonv3


type CceAddonV3Values struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/cce_addon_v3#basic CceAddonV3#basic}.
	Basic *map[string]*string `field:"required" json:"basic" yaml:"basic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/cce_addon_v3#custom CceAddonV3#custom}.
	Custom *map[string]*string `field:"required" json:"custom" yaml:"custom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/cce_addon_v3#flavor CceAddonV3#flavor}.
	Flavor *string `field:"optional" json:"flavor" yaml:"flavor"`
}

