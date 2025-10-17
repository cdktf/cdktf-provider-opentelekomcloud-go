// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedpreciseprotectionrulev1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafDedicatedPreciseProtectionRuleV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_dedicated_precise_protection_rule_v1#action WafDedicatedPreciseProtectionRuleV1#action}
	Action interface{} `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_dedicated_precise_protection_rule_v1#policy_id WafDedicatedPreciseProtectionRuleV1#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_dedicated_precise_protection_rule_v1#priority WafDedicatedPreciseProtectionRuleV1#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_dedicated_precise_protection_rule_v1#time WafDedicatedPreciseProtectionRuleV1#time}.
	Time interface{} `field:"required" json:"time" yaml:"time"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_dedicated_precise_protection_rule_v1#conditions WafDedicatedPreciseProtectionRuleV1#conditions}
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_dedicated_precise_protection_rule_v1#description WafDedicatedPreciseProtectionRuleV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_dedicated_precise_protection_rule_v1#id WafDedicatedPreciseProtectionRuleV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_dedicated_precise_protection_rule_v1#start WafDedicatedPreciseProtectionRuleV1#start}.
	Start *float64 `field:"optional" json:"start" yaml:"start"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_dedicated_precise_protection_rule_v1#terminal WafDedicatedPreciseProtectionRuleV1#terminal}.
	Terminal *float64 `field:"optional" json:"terminal" yaml:"terminal"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/waf_dedicated_precise_protection_rule_v1#timeouts WafDedicatedPreciseProtectionRuleV1#timeouts}
	Timeouts *WafDedicatedPreciseProtectionRuleV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

