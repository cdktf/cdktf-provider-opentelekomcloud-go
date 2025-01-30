// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudnatsnatrulesv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudNatSnatRulesV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/nat_snat_rules_v2#cidr DataOpentelekomcloudNatSnatRulesV2#cidr}.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/nat_snat_rules_v2#description DataOpentelekomcloudNatSnatRulesV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/nat_snat_rules_v2#floating_ip_address DataOpentelekomcloudNatSnatRulesV2#floating_ip_address}.
	FloatingIpAddress *string `field:"optional" json:"floatingIpAddress" yaml:"floatingIpAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/nat_snat_rules_v2#floating_ip_id DataOpentelekomcloudNatSnatRulesV2#floating_ip_id}.
	FloatingIpId *string `field:"optional" json:"floatingIpId" yaml:"floatingIpId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/nat_snat_rules_v2#gateway_id DataOpentelekomcloudNatSnatRulesV2#gateway_id}.
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/nat_snat_rules_v2#id DataOpentelekomcloudNatSnatRulesV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/nat_snat_rules_v2#project_id DataOpentelekomcloudNatSnatRulesV2#project_id}.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/nat_snat_rules_v2#rule_id DataOpentelekomcloudNatSnatRulesV2#rule_id}.
	RuleId *string `field:"optional" json:"ruleId" yaml:"ruleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/nat_snat_rules_v2#source_type DataOpentelekomcloudNatSnatRulesV2#source_type}.
	SourceType *float64 `field:"optional" json:"sourceType" yaml:"sourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/nat_snat_rules_v2#status DataOpentelekomcloudNatSnatRulesV2#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/data-sources/nat_snat_rules_v2#subnet_id DataOpentelekomcloudNatSnatRulesV2#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

