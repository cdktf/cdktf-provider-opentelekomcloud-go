// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwvpcchannelv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApigwVpcChannelV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/apigw_vpc_channel_v2#gateway_id ApigwVpcChannelV2#gateway_id}.
	GatewayId *string `field:"required" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/apigw_vpc_channel_v2#lb_algorithm ApigwVpcChannelV2#lb_algorithm}.
	LbAlgorithm *float64 `field:"required" json:"lbAlgorithm" yaml:"lbAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/apigw_vpc_channel_v2#name ApigwVpcChannelV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/apigw_vpc_channel_v2#port ApigwVpcChannelV2#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/apigw_vpc_channel_v2#health_check ApigwVpcChannelV2#health_check}
	HealthCheck *ApigwVpcChannelV2HealthCheck `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/apigw_vpc_channel_v2#id ApigwVpcChannelV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// member block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/apigw_vpc_channel_v2#member ApigwVpcChannelV2#member}
	Member interface{} `field:"optional" json:"member" yaml:"member"`
	// member_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/apigw_vpc_channel_v2#member_group ApigwVpcChannelV2#member_group}
	MemberGroup interface{} `field:"optional" json:"memberGroup" yaml:"memberGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/apigw_vpc_channel_v2#member_type ApigwVpcChannelV2#member_type}.
	MemberType *string `field:"optional" json:"memberType" yaml:"memberType"`
	// microservice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/apigw_vpc_channel_v2#microservice ApigwVpcChannelV2#microservice}
	Microservice *ApigwVpcChannelV2Microservice `field:"optional" json:"microservice" yaml:"microservice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/apigw_vpc_channel_v2#type ApigwVpcChannelV2#type}.
	Type *float64 `field:"optional" json:"type" yaml:"type"`
}

