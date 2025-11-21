// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedpreciseprotectionrulev1


type WafDedicatedPreciseProtectionRuleV1Action struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/waf_dedicated_precise_protection_rule_v1#category WafDedicatedPreciseProtectionRuleV1#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/waf_dedicated_precise_protection_rule_v1#followed_action_id WafDedicatedPreciseProtectionRuleV1#followed_action_id}.
	FollowedActionId *string `field:"optional" json:"followedActionId" yaml:"followedActionId"`
}

