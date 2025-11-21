// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatenatdnatrulev3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivateNatDnatRuleV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/private_nat_dnat_rule_v3#gateway_id PrivateNatDnatRuleV3#gateway_id}.
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/private_nat_dnat_rule_v3#transit_ip_id PrivateNatDnatRuleV3#transit_ip_id}.
	TransitIpId *string `field:"required" json:"transitIpId" yaml:"transitIpId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/private_nat_dnat_rule_v3#description PrivateNatDnatRuleV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/private_nat_dnat_rule_v3#internal_service_port PrivateNatDnatRuleV3#internal_service_port}.
	InternalServicePort *string `field:"optional" json:"internalServicePort" yaml:"internalServicePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/private_nat_dnat_rule_v3#network_interface_id PrivateNatDnatRuleV3#network_interface_id}.
	NetworkInterfaceId *string `field:"optional" json:"networkInterfaceId" yaml:"networkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/private_nat_dnat_rule_v3#private_ip_address PrivateNatDnatRuleV3#private_ip_address}.
	PrivateIpAddress *string `field:"optional" json:"privateIpAddress" yaml:"privateIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/private_nat_dnat_rule_v3#protocol PrivateNatDnatRuleV3#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/private_nat_dnat_rule_v3#timeouts PrivateNatDnatRuleV3#timeouts}
	Timeouts *PrivateNatDnatRuleV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/private_nat_dnat_rule_v3#transit_service_port PrivateNatDnatRuleV3#transit_service_port}.
	TransitServicePort *string `field:"optional" json:"transitServicePort" yaml:"transitServicePort"`
}

