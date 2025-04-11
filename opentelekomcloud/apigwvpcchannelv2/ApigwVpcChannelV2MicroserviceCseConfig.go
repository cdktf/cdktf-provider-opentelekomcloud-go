// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwvpcchannelv2


type ApigwVpcChannelV2MicroserviceCseConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/apigw_vpc_channel_v2#engine_id ApigwVpcChannelV2#engine_id}.
	EngineId *string `field:"required" json:"engineId" yaml:"engineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/apigw_vpc_channel_v2#service_id ApigwVpcChannelV2#service_id}.
	ServiceId *string `field:"required" json:"serviceId" yaml:"serviceId"`
}

