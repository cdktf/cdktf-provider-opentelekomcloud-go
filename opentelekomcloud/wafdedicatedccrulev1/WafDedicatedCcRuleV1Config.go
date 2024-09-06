// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedccrulev1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafDedicatedCcRuleV1Config struct {
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
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#action WafDedicatedCcRuleV1#action}
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#limit_num WafDedicatedCcRuleV1#limit_num}.
	LimitNum *float64 `field:"required" json:"limitNum" yaml:"limitNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#limit_period WafDedicatedCcRuleV1#limit_period}.
	LimitPeriod *float64 `field:"required" json:"limitPeriod" yaml:"limitPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#mode WafDedicatedCcRuleV1#mode}.
	Mode *float64 `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#policy_id WafDedicatedCcRuleV1#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#tag_type WafDedicatedCcRuleV1#tag_type}.
	TagType *string `field:"required" json:"tagType" yaml:"tagType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#url WafDedicatedCcRuleV1#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#conditions WafDedicatedCcRuleV1#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#description WafDedicatedCcRuleV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#id WafDedicatedCcRuleV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#lock_time WafDedicatedCcRuleV1#lock_time}.
	LockTime *float64 `field:"optional" json:"lockTime" yaml:"lockTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#tag_category WafDedicatedCcRuleV1#tag_category}.
	TagCategory *string `field:"optional" json:"tagCategory" yaml:"tagCategory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#tag_contents WafDedicatedCcRuleV1#tag_contents}.
	TagContents *[]*string `field:"optional" json:"tagContents" yaml:"tagContents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#tag_index WafDedicatedCcRuleV1#tag_index}.
	TagIndex *string `field:"optional" json:"tagIndex" yaml:"tagIndex"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#timeouts WafDedicatedCcRuleV1#timeouts}
	Timeouts *WafDedicatedCcRuleV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.18/docs/resources/waf_dedicated_cc_rule_v1#unlock_num WafDedicatedCcRuleV1#unlock_num}.
	UnlockNum *float64 `field:"optional" json:"unlockNum" yaml:"unlockNum"`
}

