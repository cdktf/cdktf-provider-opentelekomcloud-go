// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatenatgatewayv3


type PrivateNatGatewayV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/private_nat_gateway_v3#create PrivateNatGatewayV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/private_nat_gateway_v3#delete PrivateNatGatewayV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

