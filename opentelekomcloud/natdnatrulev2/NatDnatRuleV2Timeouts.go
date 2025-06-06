// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package natdnatrulev2


type NatDnatRuleV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/nat_dnat_rule_v2#create NatDnatRuleV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/nat_dnat_rule_v2#delete NatDnatRuleV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

