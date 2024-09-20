// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedpolicyv1


type WafDedicatedPolicyV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#create WafDedicatedPolicyV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#delete WafDedicatedPolicyV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

