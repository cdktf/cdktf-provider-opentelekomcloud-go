// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcvirtualinterfacev3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DcVirtualInterfaceV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#bandwidth DcVirtualInterfaceV3#bandwidth}.
	Bandwidth *float64 `field:"required" json:"bandwidth" yaml:"bandwidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#direct_connect_id DcVirtualInterfaceV3#direct_connect_id}.
	DirectConnectId *string `field:"required" json:"directConnectId" yaml:"directConnectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#name DcVirtualInterfaceV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#remote_ep_group DcVirtualInterfaceV3#remote_ep_group}.
	RemoteEpGroup *[]*string `field:"required" json:"remoteEpGroup" yaml:"remoteEpGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#route_mode DcVirtualInterfaceV3#route_mode}.
	RouteMode *string `field:"required" json:"routeMode" yaml:"routeMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#type DcVirtualInterfaceV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#vgw_id DcVirtualInterfaceV3#vgw_id}.
	VgwId *string `field:"required" json:"vgwId" yaml:"vgwId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#vlan DcVirtualInterfaceV3#vlan}.
	Vlan *float64 `field:"required" json:"vlan" yaml:"vlan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#address_family DcVirtualInterfaceV3#address_family}.
	AddressFamily *string `field:"optional" json:"addressFamily" yaml:"addressFamily"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#asn DcVirtualInterfaceV3#asn}.
	Asn *float64 `field:"optional" json:"asn" yaml:"asn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#bgp_md5 DcVirtualInterfaceV3#bgp_md5}.
	BgpMd5 *string `field:"optional" json:"bgpMd5" yaml:"bgpMd5"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#description DcVirtualInterfaceV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#enable_bfd DcVirtualInterfaceV3#enable_bfd}.
	EnableBfd interface{} `field:"optional" json:"enableBfd" yaml:"enableBfd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#enable_nqa DcVirtualInterfaceV3#enable_nqa}.
	EnableNqa interface{} `field:"optional" json:"enableNqa" yaml:"enableNqa"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#id DcVirtualInterfaceV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#lag_id DcVirtualInterfaceV3#lag_id}.
	LagId *string `field:"optional" json:"lagId" yaml:"lagId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#local_gateway_v4_ip DcVirtualInterfaceV3#local_gateway_v4_ip}.
	LocalGatewayV4Ip *string `field:"optional" json:"localGatewayV4Ip" yaml:"localGatewayV4Ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#local_gateway_v6_ip DcVirtualInterfaceV3#local_gateway_v6_ip}.
	LocalGatewayV6Ip *string `field:"optional" json:"localGatewayV6Ip" yaml:"localGatewayV6Ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#remote_gateway_v4_ip DcVirtualInterfaceV3#remote_gateway_v4_ip}.
	RemoteGatewayV4Ip *string `field:"optional" json:"remoteGatewayV4Ip" yaml:"remoteGatewayV4Ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#remote_gateway_v6_ip DcVirtualInterfaceV3#remote_gateway_v6_ip}.
	RemoteGatewayV6Ip *string `field:"optional" json:"remoteGatewayV6Ip" yaml:"remoteGatewayV6Ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#resource_tenant_id DcVirtualInterfaceV3#resource_tenant_id}.
	ResourceTenantId *string `field:"optional" json:"resourceTenantId" yaml:"resourceTenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/dc_virtual_interface_v3#service_ep_group DcVirtualInterfaceV3#service_ep_group}.
	ServiceEpGroup *[]*string `field:"optional" json:"serviceEpGroup" yaml:"serviceEpGroup"`
}

