// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package natdnatrulev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NatDnatRuleV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/nat_dnat_rule_v2#external_service_port NatDnatRuleV2#external_service_port}.
	ExternalServicePort *float64 `field:"required" json:"externalServicePort" yaml:"externalServicePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/nat_dnat_rule_v2#floating_ip_id NatDnatRuleV2#floating_ip_id}.
	FloatingIpId *string `field:"required" json:"floatingIpId" yaml:"floatingIpId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/nat_dnat_rule_v2#internal_service_port NatDnatRuleV2#internal_service_port}.
	InternalServicePort *float64 `field:"required" json:"internalServicePort" yaml:"internalServicePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/nat_dnat_rule_v2#nat_gateway_id NatDnatRuleV2#nat_gateway_id}.
	NatGatewayId *string `field:"required" json:"natGatewayId" yaml:"natGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/nat_dnat_rule_v2#protocol NatDnatRuleV2#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/nat_dnat_rule_v2#id NatDnatRuleV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/nat_dnat_rule_v2#port_id NatDnatRuleV2#port_id}.
	PortId *string `field:"optional" json:"portId" yaml:"portId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/nat_dnat_rule_v2#private_ip NatDnatRuleV2#private_ip}.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/nat_dnat_rule_v2#timeouts NatDnatRuleV2#timeouts}
	Timeouts *NatDnatRuleV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

