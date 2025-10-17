// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatenattransitipv3


type PrivateNatTransitIpV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/private_nat_transit_ip_v3#create PrivateNatTransitIpV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/private_nat_transit_ip_v3#delete PrivateNatTransitIpV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

