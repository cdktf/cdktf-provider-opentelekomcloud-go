// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcsubnetv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcSubnetV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#cidr VpcSubnetV1#cidr}.
	Cidr *string `field:"required" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#gateway_ip VpcSubnetV1#gateway_ip}.
	GatewayIp *string `field:"required" json:"gatewayIp" yaml:"gatewayIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#name VpcSubnetV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#vpc_id VpcSubnetV1#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#availability_zone VpcSubnetV1#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#description VpcSubnetV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#dhcp_enable VpcSubnetV1#dhcp_enable}.
	DhcpEnable interface{} `field:"optional" json:"dhcpEnable" yaml:"dhcpEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#dns_list VpcSubnetV1#dns_list}.
	DnsList *[]*string `field:"optional" json:"dnsList" yaml:"dnsList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#id VpcSubnetV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#ipv6_enable VpcSubnetV1#ipv6_enable}.
	Ipv6Enable interface{} `field:"optional" json:"ipv6Enable" yaml:"ipv6Enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#ntp_addresses VpcSubnetV1#ntp_addresses}.
	NtpAddresses *string `field:"optional" json:"ntpAddresses" yaml:"ntpAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#primary_dns VpcSubnetV1#primary_dns}.
	PrimaryDns *string `field:"optional" json:"primaryDns" yaml:"primaryDns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#region VpcSubnetV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#secondary_dns VpcSubnetV1#secondary_dns}.
	SecondaryDns *string `field:"optional" json:"secondaryDns" yaml:"secondaryDns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#tags VpcSubnetV1#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.9/docs/resources/vpc_subnet_v1#timeouts VpcSubnetV1#timeouts}
	Timeouts *VpcSubnetV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

