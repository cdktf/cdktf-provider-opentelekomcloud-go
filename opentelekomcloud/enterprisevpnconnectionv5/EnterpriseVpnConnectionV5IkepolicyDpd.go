// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisevpnconnectionv5


type EnterpriseVpnConnectionV5IkepolicyDpd struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/enterprise_vpn_connection_v5#interval EnterpriseVpnConnectionV5#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/enterprise_vpn_connection_v5#msg EnterpriseVpnConnectionV5#msg}.
	Msg *string `field:"optional" json:"msg" yaml:"msg"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/enterprise_vpn_connection_v5#timeout EnterpriseVpnConnectionV5#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

