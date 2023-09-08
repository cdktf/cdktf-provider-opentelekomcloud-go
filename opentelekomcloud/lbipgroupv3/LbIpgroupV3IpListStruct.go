// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package lbipgroupv3


type LbIpgroupV3IpListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.7/docs/resources/lb_ipgroup_v3#ip LbIpgroupV3#ip}.
	Ip *string `field:"required" json:"ip" yaml:"ip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.7/docs/resources/lb_ipgroup_v3#description LbIpgroupV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

