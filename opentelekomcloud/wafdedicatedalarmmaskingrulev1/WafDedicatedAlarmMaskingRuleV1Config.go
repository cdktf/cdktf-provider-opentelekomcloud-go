// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedalarmmaskingrulev1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafDedicatedAlarmMaskingRuleV1Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/waf_dedicated_alarm_masking_rule_v1#conditions WafDedicatedAlarmMaskingRuleV1#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/waf_dedicated_alarm_masking_rule_v1#domains WafDedicatedAlarmMaskingRuleV1#domains}.
	Domains *[]*string `field:"required" json:"domains" yaml:"domains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/waf_dedicated_alarm_masking_rule_v1#policy_id WafDedicatedAlarmMaskingRuleV1#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/waf_dedicated_alarm_masking_rule_v1#rule WafDedicatedAlarmMaskingRuleV1#rule}.
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// advanced_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/waf_dedicated_alarm_masking_rule_v1#advanced_settings WafDedicatedAlarmMaskingRuleV1#advanced_settings}
	AdvancedSettings *WafDedicatedAlarmMaskingRuleV1AdvancedSettings `field:"optional" json:"advancedSettings" yaml:"advancedSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/waf_dedicated_alarm_masking_rule_v1#description WafDedicatedAlarmMaskingRuleV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/waf_dedicated_alarm_masking_rule_v1#id WafDedicatedAlarmMaskingRuleV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.28/docs/resources/waf_dedicated_alarm_masking_rule_v1#timeouts WafDedicatedAlarmMaskingRuleV1#timeouts}
	Timeouts *WafDedicatedAlarmMaskingRuleV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

