// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisevpnconnectionv5


type EnterpriseVpnConnectionV5PolicyRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/enterprise_vpn_connection_v5#destination EnterpriseVpnConnectionV5#destination}.
	Destination *[]*string `field:"optional" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.49/docs/resources/enterprise_vpn_connection_v5#source EnterpriseVpnConnectionV5#source}.
	Source *string `field:"optional" json:"source" yaml:"source"`
}

