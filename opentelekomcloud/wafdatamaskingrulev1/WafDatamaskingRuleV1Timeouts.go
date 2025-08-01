// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdatamaskingrulev1


type WafDatamaskingRuleV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/waf_datamasking_rule_v1#create WafDatamaskingRuleV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/waf_datamasking_rule_v1#delete WafDatamaskingRuleV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

