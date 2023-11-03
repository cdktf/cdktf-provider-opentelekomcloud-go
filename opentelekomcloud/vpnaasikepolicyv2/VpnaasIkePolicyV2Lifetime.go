// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnaasikepolicyv2


type VpnaasIkePolicyV2Lifetime struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/resources/vpnaas_ike_policy_v2#units VpnaasIkePolicyV2#units}.
	Units *string `field:"optional" json:"units" yaml:"units"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/resources/vpnaas_ike_policy_v2#value VpnaasIkePolicyV2#value}.
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

