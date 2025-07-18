// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gaussdbmysqlinstancev3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GaussdbMysqlInstanceV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#flavor GaussdbMysqlInstanceV3#flavor}.
	Flavor *string `field:"required" json:"flavor" yaml:"flavor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#name GaussdbMysqlInstanceV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#password GaussdbMysqlInstanceV3#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#subnet_id GaussdbMysqlInstanceV3#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#vpc_id GaussdbMysqlInstanceV3#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#availability_zone_mode GaussdbMysqlInstanceV3#availability_zone_mode}.
	AvailabilityZoneMode *string `field:"optional" json:"availabilityZoneMode" yaml:"availabilityZoneMode"`
	// backup_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#backup_strategy GaussdbMysqlInstanceV3#backup_strategy}
	BackupStrategy *GaussdbMysqlInstanceV3BackupStrategy `field:"optional" json:"backupStrategy" yaml:"backupStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#configuration_id GaussdbMysqlInstanceV3#configuration_id}.
	ConfigurationId *string `field:"optional" json:"configurationId" yaml:"configurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#configuration_name GaussdbMysqlInstanceV3#configuration_name}.
	ConfigurationName *string `field:"optional" json:"configurationName" yaml:"configurationName"`
	// datastore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#datastore GaussdbMysqlInstanceV3#datastore}
	Datastore *GaussdbMysqlInstanceV3Datastore `field:"optional" json:"datastore" yaml:"datastore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#dedicated_resource_id GaussdbMysqlInstanceV3#dedicated_resource_id}.
	DedicatedResourceId *string `field:"optional" json:"dedicatedResourceId" yaml:"dedicatedResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#dedicated_resource_name GaussdbMysqlInstanceV3#dedicated_resource_name}.
	DedicatedResourceName *string `field:"optional" json:"dedicatedResourceName" yaml:"dedicatedResourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#id GaussdbMysqlInstanceV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#master_availability_zone GaussdbMysqlInstanceV3#master_availability_zone}.
	MasterAvailabilityZone *string `field:"optional" json:"masterAvailabilityZone" yaml:"masterAvailabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#read_replicas GaussdbMysqlInstanceV3#read_replicas}.
	ReadReplicas *float64 `field:"optional" json:"readReplicas" yaml:"readReplicas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#region GaussdbMysqlInstanceV3#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#security_group_id GaussdbMysqlInstanceV3#security_group_id}.
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#timeouts GaussdbMysqlInstanceV3#timeouts}
	Timeouts *GaussdbMysqlInstanceV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/gaussdb_mysql_instance_v3#time_zone GaussdbMysqlInstanceV3#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

