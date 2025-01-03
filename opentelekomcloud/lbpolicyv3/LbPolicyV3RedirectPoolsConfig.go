// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbpolicyv3


type LbPolicyV3RedirectPoolsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/lb_policy_v3#pool_id LbPolicyV3#pool_id}.
	PoolId *string `field:"required" json:"poolId" yaml:"poolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/lb_policy_v3#weight LbPolicyV3#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

