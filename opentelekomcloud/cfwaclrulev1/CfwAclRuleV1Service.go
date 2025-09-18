// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwaclrulev1


type CfwAclRuleV1Service struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cfw_acl_rule_v1#type CfwAclRuleV1#type}.
	Type *float64 `field:"required" json:"type" yaml:"type"`
	// custom_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cfw_acl_rule_v1#custom_service CfwAclRuleV1#custom_service}
	CustomService interface{} `field:"optional" json:"customService" yaml:"customService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cfw_acl_rule_v1#dest_port CfwAclRuleV1#dest_port}.
	DestPort *string `field:"optional" json:"destPort" yaml:"destPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cfw_acl_rule_v1#predefined_group CfwAclRuleV1#predefined_group}.
	PredefinedGroup *[]*string `field:"optional" json:"predefinedGroup" yaml:"predefinedGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cfw_acl_rule_v1#protocol CfwAclRuleV1#protocol}.
	Protocol *float64 `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cfw_acl_rule_v1#protocols CfwAclRuleV1#protocols}.
	Protocols *[]*float64 `field:"optional" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cfw_acl_rule_v1#service_group CfwAclRuleV1#service_group}.
	ServiceGroup *[]*string `field:"optional" json:"serviceGroup" yaml:"serviceGroup"`
	// service_group_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cfw_acl_rule_v1#service_group_names CfwAclRuleV1#service_group_names}
	ServiceGroupNames interface{} `field:"optional" json:"serviceGroupNames" yaml:"serviceGroupNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cfw_acl_rule_v1#service_set_id CfwAclRuleV1#service_set_id}.
	ServiceSetId *string `field:"optional" json:"serviceSetId" yaml:"serviceSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cfw_acl_rule_v1#service_set_name CfwAclRuleV1#service_set_name}.
	ServiceSetName *string `field:"optional" json:"serviceSetName" yaml:"serviceSetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cfw_acl_rule_v1#service_set_type CfwAclRuleV1#service_set_type}.
	ServiceSetType *float64 `field:"optional" json:"serviceSetType" yaml:"serviceSetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cfw_acl_rule_v1#source_port CfwAclRuleV1#source_port}.
	SourcePort *string `field:"optional" json:"sourcePort" yaml:"sourcePort"`
}

