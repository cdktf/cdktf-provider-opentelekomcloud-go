// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodepoolv3


type CceNodePoolV3RootVolume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/cce_node_pool_v3#size CceNodePoolV3#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/cce_node_pool_v3#volumetype CceNodePoolV3#volumetype}.
	Volumetype *string `field:"required" json:"volumetype" yaml:"volumetype"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/cce_node_pool_v3#extend_param CceNodePoolV3#extend_param}.
	ExtendParam *string `field:"optional" json:"extendParam" yaml:"extendParam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/cce_node_pool_v3#extend_params CceNodePoolV3#extend_params}.
	ExtendParams *map[string]*string `field:"optional" json:"extendParams" yaml:"extendParams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/cce_node_pool_v3#kms_id CceNodePoolV3#kms_id}.
	KmsId *string `field:"optional" json:"kmsId" yaml:"kmsId"`
}

