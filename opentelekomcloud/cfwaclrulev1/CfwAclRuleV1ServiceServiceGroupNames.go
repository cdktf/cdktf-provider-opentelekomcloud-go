// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwaclrulev1


type CfwAclRuleV1ServiceServiceGroupNames struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cfw_acl_rule_v1#name CfwAclRuleV1#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cfw_acl_rule_v1#protocols CfwAclRuleV1#protocols}.
	Protocols *[]*float64 `field:"optional" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cfw_acl_rule_v1#service_set_type CfwAclRuleV1#service_set_type}.
	ServiceSetType *float64 `field:"optional" json:"serviceSetType" yaml:"serviceSetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cfw_acl_rule_v1#set_id CfwAclRuleV1#set_id}.
	SetId *string `field:"optional" json:"setId" yaml:"setId"`
}

