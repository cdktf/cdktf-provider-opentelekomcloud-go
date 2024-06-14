// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cceaddonv3


type CceAddonV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/cce_addon_v3#create CceAddonV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/cce_addon_v3#delete CceAddonV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

