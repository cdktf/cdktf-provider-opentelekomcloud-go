// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsmaintenancev3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsMaintenanceV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/rds_maintenance_v3#end_time RdsMaintenanceV3#end_time}.
	EndTime *string `field:"required" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/rds_maintenance_v3#instance_id RdsMaintenanceV3#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/rds_maintenance_v3#start_time RdsMaintenanceV3#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/rds_maintenance_v3#id RdsMaintenanceV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

