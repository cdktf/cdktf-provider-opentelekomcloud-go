// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedalarmmaskingrulev1


type WafDedicatedAlarmMaskingRuleV1Conditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/waf_dedicated_alarm_masking_rule_v1#category WafDedicatedAlarmMaskingRuleV1#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/waf_dedicated_alarm_masking_rule_v1#logic_operation WafDedicatedAlarmMaskingRuleV1#logic_operation}.
	LogicOperation *string `field:"required" json:"logicOperation" yaml:"logicOperation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/waf_dedicated_alarm_masking_rule_v1#contents WafDedicatedAlarmMaskingRuleV1#contents}.
	Contents *[]*string `field:"optional" json:"contents" yaml:"contents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/waf_dedicated_alarm_masking_rule_v1#index WafDedicatedAlarmMaskingRuleV1#index}.
	Index *string `field:"optional" json:"index" yaml:"index"`
}

