// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwblacklistwhitelistrulev1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CfwBlacklistWhitelistRuleV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/cfw_blacklist_whitelist_rule_v1#address CfwBlacklistWhitelistRuleV1#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/cfw_blacklist_whitelist_rule_v1#address_type CfwBlacklistWhitelistRuleV1#address_type}.
	AddressType *float64 `field:"required" json:"addressType" yaml:"addressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/cfw_blacklist_whitelist_rule_v1#direction CfwBlacklistWhitelistRuleV1#direction}.
	Direction *float64 `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/cfw_blacklist_whitelist_rule_v1#list_type CfwBlacklistWhitelistRuleV1#list_type}.
	ListType *float64 `field:"required" json:"listType" yaml:"listType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/cfw_blacklist_whitelist_rule_v1#object_id CfwBlacklistWhitelistRuleV1#object_id}.
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/cfw_blacklist_whitelist_rule_v1#port CfwBlacklistWhitelistRuleV1#port}.
	Port *string `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/cfw_blacklist_whitelist_rule_v1#protocol CfwBlacklistWhitelistRuleV1#protocol}.
	Protocol *float64 `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/cfw_blacklist_whitelist_rule_v1#description CfwBlacklistWhitelistRuleV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/cfw_blacklist_whitelist_rule_v1#timeouts CfwBlacklistWhitelistRuleV1#timeouts}
	Timeouts *CfwBlacklistWhitelistRuleV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

