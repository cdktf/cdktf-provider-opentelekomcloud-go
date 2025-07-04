// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwaclrulev1


type CfwAclRuleV1Sequence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/cfw_acl_rule_v1#bottom CfwAclRuleV1#bottom}.
	Bottom *float64 `field:"optional" json:"bottom" yaml:"bottom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/cfw_acl_rule_v1#dest_rule_id CfwAclRuleV1#dest_rule_id}.
	DestRuleId *string `field:"optional" json:"destRuleId" yaml:"destRuleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/cfw_acl_rule_v1#top CfwAclRuleV1#top}.
	Top *float64 `field:"optional" json:"top" yaml:"top"`
}

