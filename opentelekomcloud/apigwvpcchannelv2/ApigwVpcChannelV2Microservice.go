// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwvpcchannelv2


type ApigwVpcChannelV2Microservice struct {
	// cce_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/apigw_vpc_channel_v2#cce_config ApigwVpcChannelV2#cce_config}
	CceConfig *ApigwVpcChannelV2MicroserviceCceConfig `field:"optional" json:"cceConfig" yaml:"cceConfig"`
	// cse_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/apigw_vpc_channel_v2#cse_config ApigwVpcChannelV2#cse_config}
	CseConfig *ApigwVpcChannelV2MicroserviceCseConfig `field:"optional" json:"cseConfig" yaml:"cseConfig"`
}

