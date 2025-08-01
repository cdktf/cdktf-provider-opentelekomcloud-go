// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkingsubnetv2


type NetworkingSubnetV2HostRoutes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/networking_subnet_v2#destination_cidr NetworkingSubnetV2#destination_cidr}.
	DestinationCidr *string `field:"required" json:"destinationCidr" yaml:"destinationCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/networking_subnet_v2#next_hop NetworkingSubnetV2#next_hop}.
	NextHop *string `field:"required" json:"nextHop" yaml:"nextHop"`
}

