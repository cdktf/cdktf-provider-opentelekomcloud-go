// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedblacklistrulev1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafDedicatedBlacklistRuleV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/waf_dedicated_blacklist_rule_v1#action WafDedicatedBlacklistRuleV1#action}.
	Action *float64 `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/waf_dedicated_blacklist_rule_v1#ip_address WafDedicatedBlacklistRuleV1#ip_address}.
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/waf_dedicated_blacklist_rule_v1#policy_id WafDedicatedBlacklistRuleV1#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/waf_dedicated_blacklist_rule_v1#description WafDedicatedBlacklistRuleV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/waf_dedicated_blacklist_rule_v1#followed_action_id WafDedicatedBlacklistRuleV1#followed_action_id}.
	FollowedActionId *string `field:"optional" json:"followedActionId" yaml:"followedActionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/waf_dedicated_blacklist_rule_v1#id WafDedicatedBlacklistRuleV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/waf_dedicated_blacklist_rule_v1#name WafDedicatedBlacklistRuleV1#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/waf_dedicated_blacklist_rule_v1#timeouts WafDedicatedBlacklistRuleV1#timeouts}
	Timeouts *WafDedicatedBlacklistRuleV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

