// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedpreciseprotectionrulev1


type WafDedicatedPreciseProtectionRuleV1Conditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/waf_dedicated_precise_protection_rule_v1#category WafDedicatedPreciseProtectionRuleV1#category}.
	Category *string `field:"optional" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/waf_dedicated_precise_protection_rule_v1#contents WafDedicatedPreciseProtectionRuleV1#contents}.
	Contents *[]*string `field:"optional" json:"contents" yaml:"contents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/waf_dedicated_precise_protection_rule_v1#index WafDedicatedPreciseProtectionRuleV1#index}.
	Index *string `field:"optional" json:"index" yaml:"index"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/waf_dedicated_precise_protection_rule_v1#logic_operation WafDedicatedPreciseProtectionRuleV1#logic_operation}.
	LogicOperation *string `field:"optional" json:"logicOperation" yaml:"logicOperation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/waf_dedicated_precise_protection_rule_v1#value_list_id WafDedicatedPreciseProtectionRuleV1#value_list_id}.
	ValueListId *string `field:"optional" json:"valueListId" yaml:"valueListId"`
}

