// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedccrulev1


type WafDedicatedCcRuleV1Action struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/waf_dedicated_cc_rule_v1#category WafDedicatedCcRuleV1#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/waf_dedicated_cc_rule_v1#content WafDedicatedCcRuleV1#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/waf_dedicated_cc_rule_v1#content_type WafDedicatedCcRuleV1#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
}

