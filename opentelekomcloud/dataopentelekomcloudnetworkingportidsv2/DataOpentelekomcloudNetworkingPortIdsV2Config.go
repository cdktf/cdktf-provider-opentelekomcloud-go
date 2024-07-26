// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudnetworkingportidsv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudNetworkingPortIdsV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#device_id DataOpentelekomcloudNetworkingPortIdsV2#device_id}.
	DeviceId *string `field:"optional" json:"deviceId" yaml:"deviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#device_owner DataOpentelekomcloudNetworkingPortIdsV2#device_owner}.
	DeviceOwner *string `field:"optional" json:"deviceOwner" yaml:"deviceOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#fixed_ip DataOpentelekomcloudNetworkingPortIdsV2#fixed_ip}.
	FixedIp *string `field:"optional" json:"fixedIp" yaml:"fixedIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#id DataOpentelekomcloudNetworkingPortIdsV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#mac_address DataOpentelekomcloudNetworkingPortIdsV2#mac_address}.
	MacAddress *string `field:"optional" json:"macAddress" yaml:"macAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#name DataOpentelekomcloudNetworkingPortIdsV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#network_id DataOpentelekomcloudNetworkingPortIdsV2#network_id}.
	NetworkId *string `field:"optional" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#project_id DataOpentelekomcloudNetworkingPortIdsV2#project_id}.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#region DataOpentelekomcloudNetworkingPortIdsV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#security_group_ids DataOpentelekomcloudNetworkingPortIdsV2#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#sort_direction DataOpentelekomcloudNetworkingPortIdsV2#sort_direction}.
	SortDirection *string `field:"optional" json:"sortDirection" yaml:"sortDirection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#sort_key DataOpentelekomcloudNetworkingPortIdsV2#sort_key}.
	SortKey *string `field:"optional" json:"sortKey" yaml:"sortKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#status DataOpentelekomcloudNetworkingPortIdsV2#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.15/docs/data-sources/networking_port_ids_v2#tenant_id DataOpentelekomcloudNetworkingPortIdsV2#tenant_id}.
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
}

