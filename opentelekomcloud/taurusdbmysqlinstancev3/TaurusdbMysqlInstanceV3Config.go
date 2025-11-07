// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package taurusdbmysqlinstancev3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TaurusdbMysqlInstanceV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#flavor TaurusdbMysqlInstanceV3#flavor}.
	Flavor *string `field:"required" json:"flavor" yaml:"flavor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#name TaurusdbMysqlInstanceV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#password TaurusdbMysqlInstanceV3#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#subnet_id TaurusdbMysqlInstanceV3#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#vpc_id TaurusdbMysqlInstanceV3#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#availability_zone_mode TaurusdbMysqlInstanceV3#availability_zone_mode}.
	AvailabilityZoneMode *string `field:"optional" json:"availabilityZoneMode" yaml:"availabilityZoneMode"`
	// backup_strategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#backup_strategy TaurusdbMysqlInstanceV3#backup_strategy}
	BackupStrategy *TaurusdbMysqlInstanceV3BackupStrategy `field:"optional" json:"backupStrategy" yaml:"backupStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#configuration_id TaurusdbMysqlInstanceV3#configuration_id}.
	ConfigurationId *string `field:"optional" json:"configurationId" yaml:"configurationId"`
	// datastore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#datastore TaurusdbMysqlInstanceV3#datastore}
	Datastore *TaurusdbMysqlInstanceV3Datastore `field:"optional" json:"datastore" yaml:"datastore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#dedicated_resource_id TaurusdbMysqlInstanceV3#dedicated_resource_id}.
	DedicatedResourceId *string `field:"optional" json:"dedicatedResourceId" yaml:"dedicatedResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#enterprise_project_id TaurusdbMysqlInstanceV3#enterprise_project_id}.
	EnterpriseProjectId *string `field:"optional" json:"enterpriseProjectId" yaml:"enterpriseProjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#id TaurusdbMysqlInstanceV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#master_availability_zone TaurusdbMysqlInstanceV3#master_availability_zone}.
	MasterAvailabilityZone *string `field:"optional" json:"masterAvailabilityZone" yaml:"masterAvailabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#port TaurusdbMysqlInstanceV3#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#read_replicas TaurusdbMysqlInstanceV3#read_replicas}.
	ReadReplicas *float64 `field:"optional" json:"readReplicas" yaml:"readReplicas"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#seconds_level_monitoring_enabled TaurusdbMysqlInstanceV3#seconds_level_monitoring_enabled}.
	SecondsLevelMonitoringEnabled interface{} `field:"optional" json:"secondsLevelMonitoringEnabled" yaml:"secondsLevelMonitoringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#seconds_level_monitoring_period TaurusdbMysqlInstanceV3#seconds_level_monitoring_period}.
	SecondsLevelMonitoringPeriod *float64 `field:"optional" json:"secondsLevelMonitoringPeriod" yaml:"secondsLevelMonitoringPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#security_group_id TaurusdbMysqlInstanceV3#security_group_id}.
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#table_name_case_sensitivity TaurusdbMysqlInstanceV3#table_name_case_sensitivity}.
	TableNameCaseSensitivity interface{} `field:"optional" json:"tableNameCaseSensitivity" yaml:"tableNameCaseSensitivity"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#timeouts TaurusdbMysqlInstanceV3#timeouts}
	Timeouts *TaurusdbMysqlInstanceV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#time_zone TaurusdbMysqlInstanceV3#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_instance_v3#volume_size TaurusdbMysqlInstanceV3#volume_size}.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
}

