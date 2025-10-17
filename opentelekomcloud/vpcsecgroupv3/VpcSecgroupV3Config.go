// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcsecgroupv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcSecgroupV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/vpc_secgroup_v3#name VpcSecgroupV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/vpc_secgroup_v3#description VpcSecgroupV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/vpc_secgroup_v3#enterprise_project_id VpcSecgroupV3#enterprise_project_id}.
	EnterpriseProjectId *string `field:"optional" json:"enterpriseProjectId" yaml:"enterpriseProjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/vpc_secgroup_v3#id VpcSecgroupV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/vpc_secgroup_v3#tags VpcSecgroupV3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/vpc_secgroup_v3#timeouts VpcSecgroupV3#timeouts}
	Timeouts *VpcSecgroupV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

