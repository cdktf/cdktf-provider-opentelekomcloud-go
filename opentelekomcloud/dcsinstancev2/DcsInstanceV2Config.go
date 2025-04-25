// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcsinstancev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DcsInstanceV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#availability_zones DcsInstanceV2#availability_zones}.
	AvailabilityZones *[]*string `field:"required" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#capacity DcsInstanceV2#capacity}.
	Capacity *float64 `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#engine DcsInstanceV2#engine}.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#flavor DcsInstanceV2#flavor}.
	Flavor *string `field:"required" json:"flavor" yaml:"flavor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#name DcsInstanceV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#subnet_id DcsInstanceV2#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#vpc_id DcsInstanceV2#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#access_user DcsInstanceV2#access_user}.
	AccessUser *string `field:"optional" json:"accessUser" yaml:"accessUser"`
	// backup_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#backup_policy DcsInstanceV2#backup_policy}
	BackupPolicy *DcsInstanceV2BackupPolicy `field:"optional" json:"backupPolicy" yaml:"backupPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#deleted_nodes DcsInstanceV2#deleted_nodes}.
	DeletedNodes *[]*string `field:"optional" json:"deletedNodes" yaml:"deletedNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#description DcsInstanceV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#enable_whitelist DcsInstanceV2#enable_whitelist}.
	EnableWhitelist interface{} `field:"optional" json:"enableWhitelist" yaml:"enableWhitelist"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#engine_version DcsInstanceV2#engine_version}.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#id DcsInstanceV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#maintain_begin DcsInstanceV2#maintain_begin}.
	MaintainBegin *string `field:"optional" json:"maintainBegin" yaml:"maintainBegin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#maintain_end DcsInstanceV2#maintain_end}.
	MaintainEnd *string `field:"optional" json:"maintainEnd" yaml:"maintainEnd"`
	// parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#parameters DcsInstanceV2#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#password DcsInstanceV2#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#port DcsInstanceV2#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#private_ip DcsInstanceV2#private_ip}.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#rename_commands DcsInstanceV2#rename_commands}.
	RenameCommands *map[string]*string `field:"optional" json:"renameCommands" yaml:"renameCommands"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#reserved_ips DcsInstanceV2#reserved_ips}.
	ReservedIps *[]*string `field:"optional" json:"reservedIps" yaml:"reservedIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#security_group_id DcsInstanceV2#security_group_id}.
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#ssl_enable DcsInstanceV2#ssl_enable}.
	SslEnable interface{} `field:"optional" json:"sslEnable" yaml:"sslEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#tags DcsInstanceV2#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#template_id DcsInstanceV2#template_id}.
	TemplateId *string `field:"optional" json:"templateId" yaml:"templateId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#timeouts DcsInstanceV2#timeouts}
	Timeouts *DcsInstanceV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// whitelist block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/dcs_instance_v2#whitelist DcsInstanceV2#whitelist}
	Whitelist interface{} `field:"optional" json:"whitelist" yaml:"whitelist"`
}

