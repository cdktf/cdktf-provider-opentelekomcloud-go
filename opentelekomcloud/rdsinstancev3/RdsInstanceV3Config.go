// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsinstancev3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsInstanceV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#availability_zone RdsInstanceV3#availability_zone}.
	AvailabilityZone *[]*string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#db RdsInstanceV3#db}
	Db *RdsInstanceV3Db `field:"required" json:"db" yaml:"db"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#flavor RdsInstanceV3#flavor}.
	Flavor *string `field:"required" json:"flavor" yaml:"flavor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#name RdsInstanceV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#security_group_id RdsInstanceV3#security_group_id}.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#subnet_id RdsInstanceV3#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#volume RdsInstanceV3#volume}
	Volume *RdsInstanceV3Volume `field:"required" json:"volume" yaml:"volume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#vpc_id RdsInstanceV3#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// backup_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#backup_strategy RdsInstanceV3#backup_strategy}
	BackupStrategy *RdsInstanceV3BackupStrategy `field:"optional" json:"backupStrategy" yaml:"backupStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#ha_replication_mode RdsInstanceV3#ha_replication_mode}.
	HaReplicationMode *string `field:"optional" json:"haReplicationMode" yaml:"haReplicationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#id RdsInstanceV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#lower_case_table_names RdsInstanceV3#lower_case_table_names}.
	LowerCaseTableNames *string `field:"optional" json:"lowerCaseTableNames" yaml:"lowerCaseTableNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#parameters RdsInstanceV3#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#param_group_id RdsInstanceV3#param_group_id}.
	ParamGroupId *string `field:"optional" json:"paramGroupId" yaml:"paramGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#public_ips RdsInstanceV3#public_ips}.
	PublicIps *[]*string `field:"optional" json:"publicIps" yaml:"publicIps"`
	// restore_from_backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#restore_from_backup RdsInstanceV3#restore_from_backup}
	RestoreFromBackup *RdsInstanceV3RestoreFromBackup `field:"optional" json:"restoreFromBackup" yaml:"restoreFromBackup"`
	// restore_point block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#restore_point RdsInstanceV3#restore_point}
	RestorePoint *RdsInstanceV3RestorePoint `field:"optional" json:"restorePoint" yaml:"restorePoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#ssl_enable RdsInstanceV3#ssl_enable}.
	SslEnable interface{} `field:"optional" json:"sslEnable" yaml:"sslEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#tag RdsInstanceV3#tag}.
	Tag *map[string]*string `field:"optional" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#tags RdsInstanceV3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/rds_instance_v3#timeouts RdsInstanceV3#timeouts}
	Timeouts *RdsInstanceV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

