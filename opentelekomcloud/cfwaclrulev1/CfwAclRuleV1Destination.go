// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwaclrulev1


type CfwAclRuleV1Destination struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#type CfwAclRuleV1#type}.
	Type *float64 `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#address CfwAclRuleV1#address}.
	Address *string `field:"optional" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#address_group CfwAclRuleV1#address_group}.
	AddressGroup *[]*string `field:"optional" json:"addressGroup" yaml:"addressGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#address_set_id CfwAclRuleV1#address_set_id}.
	AddressSetId *string `field:"optional" json:"addressSetId" yaml:"addressSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#address_set_name CfwAclRuleV1#address_set_name}.
	AddressSetName *string `field:"optional" json:"addressSetName" yaml:"addressSetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#address_set_type CfwAclRuleV1#address_set_type}.
	AddressSetType *float64 `field:"optional" json:"addressSetType" yaml:"addressSetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#address_type CfwAclRuleV1#address_type}.
	AddressType *float64 `field:"optional" json:"addressType" yaml:"addressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#domain_address_name CfwAclRuleV1#domain_address_name}.
	DomainAddressName *string `field:"optional" json:"domainAddressName" yaml:"domainAddressName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#domain_set_id CfwAclRuleV1#domain_set_id}.
	DomainSetId *string `field:"optional" json:"domainSetId" yaml:"domainSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#domain_set_name CfwAclRuleV1#domain_set_name}.
	DomainSetName *string `field:"optional" json:"domainSetName" yaml:"domainSetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#ip_address CfwAclRuleV1#ip_address}.
	IpAddress *[]*string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#predefined_group CfwAclRuleV1#predefined_group}.
	PredefinedGroup *[]*string `field:"optional" json:"predefinedGroup" yaml:"predefinedGroup"`
	// region_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#region_list CfwAclRuleV1#region_list}
	RegionList interface{} `field:"optional" json:"regionList" yaml:"regionList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#region_list_json CfwAclRuleV1#region_list_json}.
	RegionListJson *string `field:"optional" json:"regionListJson" yaml:"regionListJson"`
}

