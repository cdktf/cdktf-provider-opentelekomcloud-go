// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbpolicyv3


type LbPolicyV3Rules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/lb_policy_v3#compare_type LbPolicyV3#compare_type}.
	CompareType *string `field:"required" json:"compareType" yaml:"compareType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/lb_policy_v3#type LbPolicyV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/lb_policy_v3#value LbPolicyV3#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

