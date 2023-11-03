// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicateddatamaskingrulev1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafDedicatedDataMaskingRuleV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/resources/waf_dedicated_data_masking_rule_v1#category WafDedicatedDataMaskingRuleV1#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/resources/waf_dedicated_data_masking_rule_v1#name WafDedicatedDataMaskingRuleV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/resources/waf_dedicated_data_masking_rule_v1#policy_id WafDedicatedDataMaskingRuleV1#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/resources/waf_dedicated_data_masking_rule_v1#url WafDedicatedDataMaskingRuleV1#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/resources/waf_dedicated_data_masking_rule_v1#description WafDedicatedDataMaskingRuleV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/resources/waf_dedicated_data_masking_rule_v1#id WafDedicatedDataMaskingRuleV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/resources/waf_dedicated_data_masking_rule_v1#timeouts WafDedicatedDataMaskingRuleV1#timeouts}
	Timeouts *WafDedicatedDataMaskingRuleV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

