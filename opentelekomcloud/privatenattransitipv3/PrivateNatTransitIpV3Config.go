// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatenattransitipv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PrivateNatTransitIpV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/private_nat_transit_ip_v3#virsubnet_id PrivateNatTransitIpV3#virsubnet_id}.
	VirsubnetId *string `field:"required" json:"virsubnetId" yaml:"virsubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/private_nat_transit_ip_v3#enterprise_project_id PrivateNatTransitIpV3#enterprise_project_id}.
	EnterpriseProjectId *string `field:"optional" json:"enterpriseProjectId" yaml:"enterpriseProjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/private_nat_transit_ip_v3#ip_address PrivateNatTransitIpV3#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/private_nat_transit_ip_v3#tags PrivateNatTransitIpV3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/private_nat_transit_ip_v3#timeouts PrivateNatTransitIpV3#timeouts}
	Timeouts *PrivateNatTransitIpV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

