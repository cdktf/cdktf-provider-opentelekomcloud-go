// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnaassiteconnectionv2


type VpnaasSiteConnectionV2Dpd struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/vpnaas_site_connection_v2#action VpnaasSiteConnectionV2#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/vpnaas_site_connection_v2#interval VpnaasSiteConnectionV2#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/vpnaas_site_connection_v2#timeout VpnaasSiteConnectionV2#timeout}.
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

