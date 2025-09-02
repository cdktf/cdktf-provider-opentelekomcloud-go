// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cfwfirewallv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CfwFirewallV1Config struct {
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
	// charge_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/cfw_firewall_v1#charge_info CfwFirewallV1#charge_info}
	ChargeInfo *CfwFirewallV1ChargeInfo `field:"required" json:"chargeInfo" yaml:"chargeInfo"`
	// flavor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/cfw_firewall_v1#flavor CfwFirewallV1#flavor}
	Flavor *CfwFirewallV1Flavor `field:"required" json:"flavor" yaml:"flavor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/cfw_firewall_v1#name CfwFirewallV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/cfw_firewall_v1#service_type CfwFirewallV1#service_type}.
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/cfw_firewall_v1#timeouts CfwFirewallV1#timeouts}
	Timeouts *CfwFirewallV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

