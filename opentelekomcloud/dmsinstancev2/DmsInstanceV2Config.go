// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsinstancev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsInstanceV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#available_zones DmsInstanceV2#available_zones}.
	AvailableZones *[]*string `field:"required" json:"availableZones" yaml:"availableZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#engine DmsInstanceV2#engine}.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#engine_version DmsInstanceV2#engine_version}.
	EngineVersion *string `field:"required" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#name DmsInstanceV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#product_id DmsInstanceV2#product_id}.
	ProductId *string `field:"required" json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#security_group_id DmsInstanceV2#security_group_id}.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#storage_space DmsInstanceV2#storage_space}.
	StorageSpace *float64 `field:"required" json:"storageSpace" yaml:"storageSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#storage_spec_code DmsInstanceV2#storage_spec_code}.
	StorageSpecCode *string `field:"required" json:"storageSpecCode" yaml:"storageSpecCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#subnet_id DmsInstanceV2#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#vpc_id DmsInstanceV2#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#access_user DmsInstanceV2#access_user}.
	AccessUser *string `field:"optional" json:"accessUser" yaml:"accessUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#description DmsInstanceV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#disk_encrypted_enable DmsInstanceV2#disk_encrypted_enable}.
	DiskEncryptedEnable interface{} `field:"optional" json:"diskEncryptedEnable" yaml:"diskEncryptedEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#disk_encrypted_key DmsInstanceV2#disk_encrypted_key}.
	DiskEncryptedKey *string `field:"optional" json:"diskEncryptedKey" yaml:"diskEncryptedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#enable_publicip DmsInstanceV2#enable_publicip}.
	EnablePublicip interface{} `field:"optional" json:"enablePublicip" yaml:"enablePublicip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#id DmsInstanceV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#maintain_begin DmsInstanceV2#maintain_begin}.
	MaintainBegin *string `field:"optional" json:"maintainBegin" yaml:"maintainBegin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#maintain_end DmsInstanceV2#maintain_end}.
	MaintainEnd *string `field:"optional" json:"maintainEnd" yaml:"maintainEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#partition_num DmsInstanceV2#partition_num}.
	PartitionNum *float64 `field:"optional" json:"partitionNum" yaml:"partitionNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#password DmsInstanceV2#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#public_bandwidth DmsInstanceV2#public_bandwidth}.
	PublicBandwidth *float64 `field:"optional" json:"publicBandwidth" yaml:"publicBandwidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#publicip_id DmsInstanceV2#publicip_id}.
	PublicipId *[]*string `field:"optional" json:"publicipId" yaml:"publicipId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#retention_policy DmsInstanceV2#retention_policy}.
	RetentionPolicy *string `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#specification DmsInstanceV2#specification}.
	Specification *string `field:"optional" json:"specification" yaml:"specification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#tags DmsInstanceV2#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/dms_instance_v2#timeouts DmsInstanceV2#timeouts}
	Timeouts *DmsInstanceV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

