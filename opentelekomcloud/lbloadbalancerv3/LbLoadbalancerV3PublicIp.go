// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbloadbalancerv3


type LbLoadbalancerV3PublicIp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/lb_loadbalancer_v3#bandwidth_charge_mode LbLoadbalancerV3#bandwidth_charge_mode}.
	BandwidthChargeMode *string `field:"optional" json:"bandwidthChargeMode" yaml:"bandwidthChargeMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/lb_loadbalancer_v3#bandwidth_name LbLoadbalancerV3#bandwidth_name}.
	BandwidthName *string `field:"optional" json:"bandwidthName" yaml:"bandwidthName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/lb_loadbalancer_v3#bandwidth_share_type LbLoadbalancerV3#bandwidth_share_type}.
	BandwidthShareType *string `field:"optional" json:"bandwidthShareType" yaml:"bandwidthShareType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/lb_loadbalancer_v3#bandwidth_size LbLoadbalancerV3#bandwidth_size}.
	BandwidthSize *float64 `field:"optional" json:"bandwidthSize" yaml:"bandwidthSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/lb_loadbalancer_v3#id LbLoadbalancerV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/lb_loadbalancer_v3#ip_type LbLoadbalancerV3#ip_type}.
	IpType *string `field:"optional" json:"ipType" yaml:"ipType"`
}

