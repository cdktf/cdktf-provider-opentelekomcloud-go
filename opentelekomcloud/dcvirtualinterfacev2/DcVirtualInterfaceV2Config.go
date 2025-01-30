// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcvirtualinterfacev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DcVirtualInterfaceV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#bandwidth DcVirtualInterfaceV2#bandwidth}.
	Bandwidth *float64 `field:"required" json:"bandwidth" yaml:"bandwidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#direct_connect_id DcVirtualInterfaceV2#direct_connect_id}.
	DirectConnectId *string `field:"required" json:"directConnectId" yaml:"directConnectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#name DcVirtualInterfaceV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// remote_ep_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#remote_ep_group DcVirtualInterfaceV2#remote_ep_group}
	RemoteEpGroup *DcVirtualInterfaceV2RemoteEpGroup `field:"required" json:"remoteEpGroup" yaml:"remoteEpGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#route_mode DcVirtualInterfaceV2#route_mode}.
	RouteMode *string `field:"required" json:"routeMode" yaml:"routeMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#service_type DcVirtualInterfaceV2#service_type}.
	ServiceType *string `field:"required" json:"serviceType" yaml:"serviceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#type DcVirtualInterfaceV2#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#virtual_gateway_id DcVirtualInterfaceV2#virtual_gateway_id}.
	VirtualGatewayId *string `field:"required" json:"virtualGatewayId" yaml:"virtualGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#vlan DcVirtualInterfaceV2#vlan}.
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#asn DcVirtualInterfaceV2#asn}.
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#bgp_md5 DcVirtualInterfaceV2#bgp_md5}.
	BgpMd5 *string `field:"optional" json:"bgpMd5" yaml:"bgpMd5"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#description DcVirtualInterfaceV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#enable_bfd DcVirtualInterfaceV2#enable_bfd}.
	EnableBfd interface{} `field:"optional" json:"enableBfd" yaml:"enableBfd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#enable_nqa DcVirtualInterfaceV2#enable_nqa}.
	EnableNqa interface{} `field:"optional" json:"enableNqa" yaml:"enableNqa"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#id DcVirtualInterfaceV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#lag_id DcVirtualInterfaceV2#lag_id}.
	LagId *string `field:"optional" json:"lagId" yaml:"lagId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#local_gateway_v4_ip DcVirtualInterfaceV2#local_gateway_v4_ip}.
	LocalGatewayV4Ip *string `field:"optional" json:"localGatewayV4Ip" yaml:"localGatewayV4Ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#project_id DcVirtualInterfaceV2#project_id}.
	ProjectId *string `field:"optional" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/dc_virtual_interface_v2#remote_gateway_v4_ip DcVirtualInterfaceV2#remote_gateway_v4_ip}.
	RemoteGatewayV4Ip *string `field:"optional" json:"remoteGatewayV4Ip" yaml:"remoteGatewayV4Ip"`
}

