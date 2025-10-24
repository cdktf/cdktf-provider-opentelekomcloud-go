// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisevpngatewayv5


type EnterpriseVpnGatewayV5Eip2 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_gateway_v5#bandwidth_name EnterpriseVpnGatewayV5#bandwidth_name}.
	BandwidthName *string `field:"optional" json:"bandwidthName" yaml:"bandwidthName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_gateway_v5#bandwidth_size EnterpriseVpnGatewayV5#bandwidth_size}.
	BandwidthSize *float64 `field:"optional" json:"bandwidthSize" yaml:"bandwidthSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_gateway_v5#charge_mode EnterpriseVpnGatewayV5#charge_mode}.
	ChargeMode *string `field:"optional" json:"chargeMode" yaml:"chargeMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_gateway_v5#id EnterpriseVpnGatewayV5#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/enterprise_vpn_gateway_v5#type EnterpriseVpnGatewayV5#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

