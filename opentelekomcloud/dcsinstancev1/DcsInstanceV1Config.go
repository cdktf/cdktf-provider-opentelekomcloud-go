// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcsinstancev1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DcsInstanceV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#available_zones DcsInstanceV1#available_zones}.
	AvailableZones *[]*string `field:"required" json:"availableZones" yaml:"availableZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#capacity DcsInstanceV1#capacity}.
	Capacity *float64 `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#engine DcsInstanceV1#engine}.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#engine_version DcsInstanceV1#engine_version}.
	EngineVersion *string `field:"required" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#name DcsInstanceV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#product_id DcsInstanceV1#product_id}.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#subnet_id DcsInstanceV1#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#vpc_id DcsInstanceV1#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#backup_at DcsInstanceV1#backup_at}.
	BackupAt *[]*float64 `field:"optional" json:"backupAt" yaml:"backupAt"`
	// backup_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#backup_policy DcsInstanceV1#backup_policy}
	BackupPolicy *DcsInstanceV1BackupPolicy `field:"optional" json:"backupPolicy" yaml:"backupPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#backup_type DcsInstanceV1#backup_type}.
	BackupType *string `field:"optional" json:"backupType" yaml:"backupType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#begin_at DcsInstanceV1#begin_at}.
	BeginAt *string `field:"optional" json:"beginAt" yaml:"beginAt"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#configuration DcsInstanceV1#configuration}
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#description DcsInstanceV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#enable_whitelist DcsInstanceV1#enable_whitelist}.
	EnableWhitelist interface{} `field:"optional" json:"enableWhitelist" yaml:"enableWhitelist"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#id DcsInstanceV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#maintain_begin DcsInstanceV1#maintain_begin}.
	MaintainBegin *string `field:"optional" json:"maintainBegin" yaml:"maintainBegin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#maintain_end DcsInstanceV1#maintain_end}.
	MaintainEnd *string `field:"optional" json:"maintainEnd" yaml:"maintainEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#password DcsInstanceV1#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#period_type DcsInstanceV1#period_type}.
	PeriodType *string `field:"optional" json:"periodType" yaml:"periodType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#private_ip DcsInstanceV1#private_ip}.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#save_days DcsInstanceV1#save_days}.
	SaveDays *float64 `field:"optional" json:"saveDays" yaml:"saveDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#security_group_id DcsInstanceV1#security_group_id}.
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#tags DcsInstanceV1#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#timeouts DcsInstanceV1#timeouts}
	Timeouts *DcsInstanceV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// whitelist block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.31/docs/resources/dcs_instance_v1#whitelist DcsInstanceV1#whitelist}
	Whitelist interface{} `field:"optional" json:"whitelist" yaml:"whitelist"`
}

