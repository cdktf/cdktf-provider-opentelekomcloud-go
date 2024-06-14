// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package natsnatrulev2


type NatSnatRuleV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/nat_snat_rule_v2#create NatSnatRuleV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/nat_snat_rule_v2#delete NatSnatRuleV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

