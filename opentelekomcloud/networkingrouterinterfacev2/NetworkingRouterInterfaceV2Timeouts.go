// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkingrouterinterfacev2


type NetworkingRouterInterfaceV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/networking_router_interface_v2#create NetworkingRouterInterfaceV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/networking_router_interface_v2#delete NetworkingRouterInterfaceV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

