// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatenatgatewayv3


type PrivateNatGatewayV3DownlinkVpcs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/private_nat_gateway_v3#virsubnet_id PrivateNatGatewayV3#virsubnet_id}.
	VirsubnetId *string `field:"required" json:"virsubnetId" yaml:"virsubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/private_nat_gateway_v3#ngport_ip_address PrivateNatGatewayV3#ngport_ip_address}.
	NgportIpAddress *string `field:"optional" json:"ngportIpAddress" yaml:"ngportIpAddress"`
}

