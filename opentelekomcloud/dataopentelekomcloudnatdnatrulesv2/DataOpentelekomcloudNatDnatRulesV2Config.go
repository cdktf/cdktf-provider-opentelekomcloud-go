// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudnatdnatrulesv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudNatDnatRulesV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/nat_dnat_rules_v2#description DataOpentelekomcloudNatDnatRulesV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/nat_dnat_rules_v2#external_service_port DataOpentelekomcloudNatDnatRulesV2#external_service_port}.
	ExternalServicePort *float64 `field:"optional" json:"externalServicePort" yaml:"externalServicePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/nat_dnat_rules_v2#floating_ip_address DataOpentelekomcloudNatDnatRulesV2#floating_ip_address}.
	FloatingIpAddress *string `field:"optional" json:"floatingIpAddress" yaml:"floatingIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/nat_dnat_rules_v2#floating_ip_id DataOpentelekomcloudNatDnatRulesV2#floating_ip_id}.
	FloatingIpId *string `field:"optional" json:"floatingIpId" yaml:"floatingIpId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/nat_dnat_rules_v2#gateway_id DataOpentelekomcloudNatDnatRulesV2#gateway_id}.
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/nat_dnat_rules_v2#id DataOpentelekomcloudNatDnatRulesV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/nat_dnat_rules_v2#internal_service_port DataOpentelekomcloudNatDnatRulesV2#internal_service_port}.
	InternalServicePort *float64 `field:"optional" json:"internalServicePort" yaml:"internalServicePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/nat_dnat_rules_v2#port_id DataOpentelekomcloudNatDnatRulesV2#port_id}.
	PortId *string `field:"optional" json:"portId" yaml:"portId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/nat_dnat_rules_v2#private_ip DataOpentelekomcloudNatDnatRulesV2#private_ip}.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/nat_dnat_rules_v2#protocol DataOpentelekomcloudNatDnatRulesV2#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/nat_dnat_rules_v2#rule_id DataOpentelekomcloudNatDnatRulesV2#rule_id}.
	RuleId *string `field:"optional" json:"ruleId" yaml:"ruleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/nat_dnat_rules_v2#status DataOpentelekomcloudNatDnatRulesV2#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

