// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedccrulev1


type WafDedicatedCcRuleV1Conditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/waf_dedicated_cc_rule_v1#category WafDedicatedCcRuleV1#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/waf_dedicated_cc_rule_v1#logic_operation WafDedicatedCcRuleV1#logic_operation}.
	LogicOperation *string `field:"required" json:"logicOperation" yaml:"logicOperation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/waf_dedicated_cc_rule_v1#contents WafDedicatedCcRuleV1#contents}.
	Contents *[]*string `field:"optional" json:"contents" yaml:"contents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/waf_dedicated_cc_rule_v1#index WafDedicatedCcRuleV1#index}.
	Index *string `field:"optional" json:"index" yaml:"index"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/waf_dedicated_cc_rule_v1#value_list_id WafDedicatedCcRuleV1#value_list_id}.
	ValueListId *string `field:"optional" json:"valueListId" yaml:"valueListId"`
}

