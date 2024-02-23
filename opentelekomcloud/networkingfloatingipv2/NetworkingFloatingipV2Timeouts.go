// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkingfloatingipv2


type NetworkingFloatingipV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/networking_floatingip_v2#create NetworkingFloatingipV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/networking_floatingip_v2#delete NetworkingFloatingipV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

