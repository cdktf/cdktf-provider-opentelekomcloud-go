// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ddsinstancev3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DdsInstanceV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#availability_zone DdsInstanceV3#availability_zone}.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// datastore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#datastore DdsInstanceV3#datastore}
	Datastore *DdsInstanceV3Datastore `field:"required" json:"datastore" yaml:"datastore"`
	// flavor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#flavor DdsInstanceV3#flavor}
	Flavor interface{} `field:"required" json:"flavor" yaml:"flavor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#mode DdsInstanceV3#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#name DdsInstanceV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#password DdsInstanceV3#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#security_group_id DdsInstanceV3#security_group_id}.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#subnet_id DdsInstanceV3#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#vpc_id DdsInstanceV3#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// backup_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#backup_strategy DdsInstanceV3#backup_strategy}
	BackupStrategy *DdsInstanceV3BackupStrategy `field:"optional" json:"backupStrategy" yaml:"backupStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#disk_encryption_id DdsInstanceV3#disk_encryption_id}.
	DiskEncryptionId *string `field:"optional" json:"diskEncryptionId" yaml:"diskEncryptionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#id DdsInstanceV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#port DdsInstanceV3#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#region DdsInstanceV3#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#ssl DdsInstanceV3#ssl}.
	Ssl interface{} `field:"optional" json:"ssl" yaml:"ssl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#tags DdsInstanceV3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dds_instance_v3#timeouts DdsInstanceV3#timeouts}
	Timeouts *DdsInstanceV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

