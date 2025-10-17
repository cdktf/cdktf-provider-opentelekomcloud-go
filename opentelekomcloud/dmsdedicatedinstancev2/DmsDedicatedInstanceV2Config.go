// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dmsdedicatedinstancev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DmsDedicatedInstanceV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#broker_num DmsDedicatedInstanceV2#broker_num}.
	BrokerNum *float64 `field:"required" json:"brokerNum" yaml:"brokerNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#engine_version DmsDedicatedInstanceV2#engine_version}.
	EngineVersion *string `field:"required" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#flavor_id DmsDedicatedInstanceV2#flavor_id}.
	FlavorId *string `field:"required" json:"flavorId" yaml:"flavorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#name DmsDedicatedInstanceV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#network_id DmsDedicatedInstanceV2#network_id}.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#security_group_id DmsDedicatedInstanceV2#security_group_id}.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#storage_space DmsDedicatedInstanceV2#storage_space}.
	StorageSpace *float64 `field:"required" json:"storageSpace" yaml:"storageSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#storage_spec_code DmsDedicatedInstanceV2#storage_spec_code}.
	StorageSpecCode *string `field:"required" json:"storageSpecCode" yaml:"storageSpecCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#vpc_id DmsDedicatedInstanceV2#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#access_user DmsDedicatedInstanceV2#access_user}.
	AccessUser *string `field:"optional" json:"accessUser" yaml:"accessUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#arch_type DmsDedicatedInstanceV2#arch_type}.
	ArchType *string `field:"optional" json:"archType" yaml:"archType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#available_zones DmsDedicatedInstanceV2#available_zones}.
	AvailableZones *[]*string `field:"optional" json:"availableZones" yaml:"availableZones"`
	// cross_vpc_accesses block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#cross_vpc_accesses DmsDedicatedInstanceV2#cross_vpc_accesses}
	CrossVpcAccesses interface{} `field:"optional" json:"crossVpcAccesses" yaml:"crossVpcAccesses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#description DmsDedicatedInstanceV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#disk_encrypted_enable DmsDedicatedInstanceV2#disk_encrypted_enable}.
	DiskEncryptedEnable interface{} `field:"optional" json:"diskEncryptedEnable" yaml:"diskEncryptedEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#disk_encrypted_key DmsDedicatedInstanceV2#disk_encrypted_key}.
	DiskEncryptedKey *string `field:"optional" json:"diskEncryptedKey" yaml:"diskEncryptedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#enabled_mechanisms DmsDedicatedInstanceV2#enabled_mechanisms}.
	EnabledMechanisms *[]*string `field:"optional" json:"enabledMechanisms" yaml:"enabledMechanisms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#enable_publicip DmsDedicatedInstanceV2#enable_publicip}.
	EnablePublicip interface{} `field:"optional" json:"enablePublicip" yaml:"enablePublicip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#id DmsDedicatedInstanceV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#ipv6_enable DmsDedicatedInstanceV2#ipv6_enable}.
	Ipv6Enable interface{} `field:"optional" json:"ipv6Enable" yaml:"ipv6Enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#maintain_begin DmsDedicatedInstanceV2#maintain_begin}.
	MaintainBegin *string `field:"optional" json:"maintainBegin" yaml:"maintainBegin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#maintain_end DmsDedicatedInstanceV2#maintain_end}.
	MaintainEnd *string `field:"optional" json:"maintainEnd" yaml:"maintainEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#new_tenant_ips DmsDedicatedInstanceV2#new_tenant_ips}.
	NewTenantIps *[]*string `field:"optional" json:"newTenantIps" yaml:"newTenantIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#password DmsDedicatedInstanceV2#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#publicip_id DmsDedicatedInstanceV2#publicip_id}.
	PublicipId *[]*string `field:"optional" json:"publicipId" yaml:"publicipId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#retention_policy DmsDedicatedInstanceV2#retention_policy}.
	RetentionPolicy *string `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#security_protocol DmsDedicatedInstanceV2#security_protocol}.
	SecurityProtocol *string `field:"optional" json:"securityProtocol" yaml:"securityProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#ssl_enable DmsDedicatedInstanceV2#ssl_enable}.
	SslEnable interface{} `field:"optional" json:"sslEnable" yaml:"sslEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#tags DmsDedicatedInstanceV2#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/dms_dedicated_instance_v2#timeouts DmsDedicatedInstanceV2#timeouts}
	Timeouts *DmsDedicatedInstanceV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

