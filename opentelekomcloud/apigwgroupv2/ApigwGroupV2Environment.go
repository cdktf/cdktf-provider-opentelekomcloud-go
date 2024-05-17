// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwgroupv2


type ApigwGroupV2Environment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/apigw_group_v2#environment_id ApigwGroupV2#environment_id}.
	EnvironmentId *string `field:"required" json:"environmentId" yaml:"environmentId"`
	// variable block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.8/docs/resources/apigw_group_v2#variable ApigwGroupV2#variable}
	Variable interface{} `field:"required" json:"variable" yaml:"variable"`
}

