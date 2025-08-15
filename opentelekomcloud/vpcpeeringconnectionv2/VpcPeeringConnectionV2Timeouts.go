// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcpeeringconnectionv2


type VpcPeeringConnectionV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/vpc_peering_connection_v2#create VpcPeeringConnectionV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/vpc_peering_connection_v2#delete VpcPeeringConnectionV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

