// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dcsinstancev2


type DcsInstanceV2WhitelistStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dcs_instance_v2#group_name DcsInstanceV2#group_name}.
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/dcs_instance_v2#ip_list DcsInstanceV2#ip_list}.
	IpList *[]*string `field:"required" json:"ipList" yaml:"ipList"`
}

