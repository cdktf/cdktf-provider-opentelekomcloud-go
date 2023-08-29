// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnaasipsecpolicyv2


type VpnaasIpsecPolicyV2Lifetime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vpnaas_ipsec_policy_v2#units VpnaasIpsecPolicyV2#units}.
	Units *string `field:"optional" json:"units" yaml:"units"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vpnaas_ipsec_policy_v2#value VpnaasIpsecPolicyV2#value}.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

