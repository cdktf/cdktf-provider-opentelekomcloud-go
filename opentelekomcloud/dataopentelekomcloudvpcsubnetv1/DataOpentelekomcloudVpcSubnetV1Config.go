// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudvpcsubnetv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudVpcSubnetV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/data-sources/vpc_subnet_v1#availability_zone DataOpentelekomcloudVpcSubnetV1#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/data-sources/vpc_subnet_v1#cidr DataOpentelekomcloudVpcSubnetV1#cidr}.
	Cidr *string `field:"optional" json:"cidr" yaml:"cidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/data-sources/vpc_subnet_v1#gateway_ip DataOpentelekomcloudVpcSubnetV1#gateway_ip}.
	GatewayIp *string `field:"optional" json:"gatewayIp" yaml:"gatewayIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/data-sources/vpc_subnet_v1#id DataOpentelekomcloudVpcSubnetV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/data-sources/vpc_subnet_v1#name DataOpentelekomcloudVpcSubnetV1#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/data-sources/vpc_subnet_v1#primary_dns DataOpentelekomcloudVpcSubnetV1#primary_dns}.
	PrimaryDns *string `field:"optional" json:"primaryDns" yaml:"primaryDns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/data-sources/vpc_subnet_v1#region DataOpentelekomcloudVpcSubnetV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/data-sources/vpc_subnet_v1#secondary_dns DataOpentelekomcloudVpcSubnetV1#secondary_dns}.
	SecondaryDns *string `field:"optional" json:"secondaryDns" yaml:"secondaryDns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/data-sources/vpc_subnet_v1#status DataOpentelekomcloudVpcSubnetV1#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/data-sources/vpc_subnet_v1#vpc_id DataOpentelekomcloudVpcSubnetV1#vpc_id}.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

