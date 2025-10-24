// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbl7rulev2


type LbL7RuleV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/lb_l7rule_v2#create LbL7RuleV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/lb_l7rule_v2#delete LbL7RuleV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/lb_l7rule_v2#update LbL7RuleV2#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

