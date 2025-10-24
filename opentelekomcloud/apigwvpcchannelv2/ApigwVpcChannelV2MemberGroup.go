// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwvpcchannelv2


type ApigwVpcChannelV2MemberGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/apigw_vpc_channel_v2#name ApigwVpcChannelV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/apigw_vpc_channel_v2#description ApigwVpcChannelV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/apigw_vpc_channel_v2#microservice_port ApigwVpcChannelV2#microservice_port}.
	MicroservicePort *float64 `field:"optional" json:"microservicePort" yaml:"microservicePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/apigw_vpc_channel_v2#microservice_tags ApigwVpcChannelV2#microservice_tags}.
	MicroserviceTags *map[string]*string `field:"optional" json:"microserviceTags" yaml:"microserviceTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/apigw_vpc_channel_v2#microservice_version ApigwVpcChannelV2#microservice_version}.
	MicroserviceVersion *string `field:"optional" json:"microserviceVersion" yaml:"microserviceVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/apigw_vpc_channel_v2#weight ApigwVpcChannelV2#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

