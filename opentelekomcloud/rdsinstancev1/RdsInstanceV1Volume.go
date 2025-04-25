// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsinstancev1


type RdsInstanceV1Volume struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/rds_instance_v1#size RdsInstanceV1#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/rds_instance_v1#type RdsInstanceV1#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

