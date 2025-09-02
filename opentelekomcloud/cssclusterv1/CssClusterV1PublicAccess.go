// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cssclusterv1


type CssClusterV1PublicAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/css_cluster_v1#bandwidth CssClusterV1#bandwidth}.
	Bandwidth *float64 `field:"required" json:"bandwidth" yaml:"bandwidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/css_cluster_v1#whitelist_enabled CssClusterV1#whitelist_enabled}.
	WhitelistEnabled interface{} `field:"required" json:"whitelistEnabled" yaml:"whitelistEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/css_cluster_v1#whitelist CssClusterV1#whitelist}.
	Whitelist *string `field:"optional" json:"whitelist" yaml:"whitelist"`
}

