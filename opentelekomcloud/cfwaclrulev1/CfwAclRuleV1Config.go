// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwaclrulev1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CfwAclRuleV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#action_type CfwAclRuleV1#action_type}.
	ActionType *float64 `field:"required" json:"actionType" yaml:"actionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#address_type CfwAclRuleV1#address_type}.
	AddressType *float64 `field:"required" json:"addressType" yaml:"addressType"`
	// destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#destination CfwAclRuleV1#destination}
	Destination *CfwAclRuleV1Destination `field:"required" json:"destination" yaml:"destination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#long_connect_enable CfwAclRuleV1#long_connect_enable}.
	LongConnectEnable *float64 `field:"required" json:"longConnectEnable" yaml:"longConnectEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#name CfwAclRuleV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#object_id CfwAclRuleV1#object_id}.
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// sequence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#sequence CfwAclRuleV1#sequence}
	Sequence *CfwAclRuleV1Sequence `field:"required" json:"sequence" yaml:"sequence"`
	// service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#service CfwAclRuleV1#service}
	Service *CfwAclRuleV1Service `field:"required" json:"service" yaml:"service"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#source CfwAclRuleV1#source}
	Source *CfwAclRuleV1Source `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#status CfwAclRuleV1#status}.
	Status *float64 `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#type CfwAclRuleV1#type}.
	Type *float64 `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#applications CfwAclRuleV1#applications}.
	Applications *[]*string `field:"optional" json:"applications" yaml:"applications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#applications_json_string CfwAclRuleV1#applications_json_string}.
	ApplicationsJsonString *string `field:"optional" json:"applicationsJsonString" yaml:"applicationsJsonString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#description CfwAclRuleV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#direction CfwAclRuleV1#direction}.
	Direction *float64 `field:"optional" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#long_connect_time CfwAclRuleV1#long_connect_time}.
	LongConnectTime *float64 `field:"optional" json:"longConnectTime" yaml:"longConnectTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#long_connect_time_hour CfwAclRuleV1#long_connect_time_hour}.
	LongConnectTimeHour *float64 `field:"optional" json:"longConnectTimeHour" yaml:"longConnectTimeHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#long_connect_time_minute CfwAclRuleV1#long_connect_time_minute}.
	LongConnectTimeMinute *float64 `field:"optional" json:"longConnectTimeMinute" yaml:"longConnectTimeMinute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#long_connect_time_second CfwAclRuleV1#long_connect_time_second}.
	LongConnectTimeSecond *float64 `field:"optional" json:"longConnectTimeSecond" yaml:"longConnectTimeSecond"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/cfw_acl_rule_v1#timeouts CfwAclRuleV1#timeouts}
	Timeouts *CfwAclRuleV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

