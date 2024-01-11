// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbpoolv2


type LbPoolV2Persistence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/lb_pool_v2#cookie_name LbPoolV2#cookie_name}.
	CookieName *string `field:"optional" json:"cookieName" yaml:"cookieName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/lb_pool_v2#type LbPoolV2#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

