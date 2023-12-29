// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package natsnatrulev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NatSnatRuleV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/nat_snat_rule_v2#floating_ip_id NatSnatRuleV2#floating_ip_id}.
	FloatingIpId *string `field:"required" json:"floatingIpId" yaml:"floatingIpId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/nat_snat_rule_v2#nat_gateway_id NatSnatRuleV2#nat_gateway_id}.
	NatGatewayId *string `field:"required" json:"natGatewayId" yaml:"natGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/nat_snat_rule_v2#cidr NatSnatRuleV2#cidr}.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/nat_snat_rule_v2#id NatSnatRuleV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/nat_snat_rule_v2#network_id NatSnatRuleV2#network_id}.
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/nat_snat_rule_v2#region NatSnatRuleV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/nat_snat_rule_v2#source_type NatSnatRuleV2#source_type}.
	SourceType *float64 `field:"optional" json:"sourceType" yaml:"sourceType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/nat_snat_rule_v2#timeouts NatSnatRuleV2#timeouts}
	Timeouts *NatSnatRuleV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

