// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcsecgrouprulev3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcSecgroupRuleV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/vpc_secgroup_rule_v3#direction VpcSecgroupRuleV3#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/vpc_secgroup_rule_v3#security_group_id VpcSecgroupRuleV3#security_group_id}.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/vpc_secgroup_rule_v3#action VpcSecgroupRuleV3#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/vpc_secgroup_rule_v3#description VpcSecgroupRuleV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/vpc_secgroup_rule_v3#ether_type VpcSecgroupRuleV3#ether_type}.
	EtherType *string `field:"optional" json:"etherType" yaml:"etherType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/vpc_secgroup_rule_v3#id VpcSecgroupRuleV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/vpc_secgroup_rule_v3#multi_port VpcSecgroupRuleV3#multi_port}.
	MultiPort *string `field:"optional" json:"multiPort" yaml:"multiPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/vpc_secgroup_rule_v3#priority VpcSecgroupRuleV3#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/vpc_secgroup_rule_v3#protocol VpcSecgroupRuleV3#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/vpc_secgroup_rule_v3#remote_group_id VpcSecgroupRuleV3#remote_group_id}.
	RemoteGroupId *string `field:"optional" json:"remoteGroupId" yaml:"remoteGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/vpc_secgroup_rule_v3#remote_ip_prefix VpcSecgroupRuleV3#remote_ip_prefix}.
	RemoteIpPrefix *string `field:"optional" json:"remoteIpPrefix" yaml:"remoteIpPrefix"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/vpc_secgroup_rule_v3#timeouts VpcSecgroupRuleV3#timeouts}
	Timeouts *VpcSecgroupRuleV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

