// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwvpcchannelv2


type ApigwVpcChannelV2Member struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/apigw_vpc_channel_v2#group_name ApigwVpcChannelV2#group_name}.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/apigw_vpc_channel_v2#host ApigwVpcChannelV2#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/apigw_vpc_channel_v2#id ApigwVpcChannelV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/apigw_vpc_channel_v2#is_backup ApigwVpcChannelV2#is_backup}.
	IsBackup interface{} `field:"optional" json:"isBackup" yaml:"isBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/apigw_vpc_channel_v2#name ApigwVpcChannelV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/apigw_vpc_channel_v2#port ApigwVpcChannelV2#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/apigw_vpc_channel_v2#status ApigwVpcChannelV2#status}.
	Status *float64 `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.40/docs/resources/apigw_vpc_channel_v2#weight ApigwVpcChannelV2#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

