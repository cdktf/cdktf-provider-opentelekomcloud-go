// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwgroupv2


type ApigwGroupV2EnvironmentVariable struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/apigw_group_v2#name ApigwGroupV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/apigw_group_v2#value ApigwGroupV2#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

