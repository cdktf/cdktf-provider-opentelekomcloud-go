// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkingsubnetv2


type NetworkingSubnetV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.17/docs/resources/networking_subnet_v2#create NetworkingSubnetV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.17/docs/resources/networking_subnet_v2#delete NetworkingSubnetV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

