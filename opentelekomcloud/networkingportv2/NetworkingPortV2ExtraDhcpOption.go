// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkingportv2


type NetworkingPortV2ExtraDhcpOption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/networking_port_v2#name NetworkingPortV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/networking_port_v2#value NetworkingPortV2#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

