// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwdomainnamegroupv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CfwDomainNameGroupV1Config struct {
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
	// domain_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/cfw_domain_name_group_v1#domain_names CfwDomainNameGroupV1#domain_names}
	DomainNames interface{} `field:"required" json:"domainNames" yaml:"domainNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/cfw_domain_name_group_v1#firewall_id CfwDomainNameGroupV1#firewall_id}.
	FirewallId *string `field:"required" json:"firewallId" yaml:"firewallId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/cfw_domain_name_group_v1#name CfwDomainNameGroupV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/cfw_domain_name_group_v1#object_id CfwDomainNameGroupV1#object_id}.
	ObjectId *string `field:"required" json:"objectId" yaml:"objectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/cfw_domain_name_group_v1#description CfwDomainNameGroupV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/cfw_domain_name_group_v1#domain_set_type CfwDomainNameGroupV1#domain_set_type}.
	DomainSetType *float64 `field:"optional" json:"domainSetType" yaml:"domainSetType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/cfw_domain_name_group_v1#timeouts CfwDomainNameGroupV1#timeouts}
	Timeouts *CfwDomainNameGroupV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

