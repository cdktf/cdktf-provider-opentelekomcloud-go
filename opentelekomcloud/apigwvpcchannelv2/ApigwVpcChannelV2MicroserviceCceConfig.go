// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwvpcchannelv2


type ApigwVpcChannelV2MicroserviceCceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_vpc_channel_v2#cluster_id ApigwVpcChannelV2#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_vpc_channel_v2#namespace ApigwVpcChannelV2#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_vpc_channel_v2#workload_type ApigwVpcChannelV2#workload_type}.
	WorkloadType *string `field:"required" json:"workloadType" yaml:"workloadType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_vpc_channel_v2#label_key ApigwVpcChannelV2#label_key}.
	LabelKey *string `field:"optional" json:"labelKey" yaml:"labelKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_vpc_channel_v2#label_value ApigwVpcChannelV2#label_value}.
	LabelValue *string `field:"optional" json:"labelValue" yaml:"labelValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_vpc_channel_v2#workload_name ApigwVpcChannelV2#workload_name}.
	WorkloadName *string `field:"optional" json:"workloadName" yaml:"workloadName"`
}

