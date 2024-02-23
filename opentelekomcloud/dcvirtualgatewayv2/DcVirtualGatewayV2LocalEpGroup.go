// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcvirtualgatewayv2


type DcVirtualGatewayV2LocalEpGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/dc_virtual_gateway_v2#endpoints DcVirtualGatewayV2#endpoints}.
	Endpoints *[]*string `field:"required" json:"endpoints" yaml:"endpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/dc_virtual_gateway_v2#description DcVirtualGatewayV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/dc_virtual_gateway_v2#name DcVirtualGatewayV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.2/docs/resources/dc_virtual_gateway_v2#type DcVirtualGatewayV2#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

