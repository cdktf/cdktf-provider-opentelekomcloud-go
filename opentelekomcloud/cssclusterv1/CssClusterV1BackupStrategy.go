// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cssclusterv1


type CssClusterV1BackupStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/css_cluster_v1#keep_days CssClusterV1#keep_days}.
	KeepDays *float64 `field:"required" json:"keepDays" yaml:"keepDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/css_cluster_v1#prefix CssClusterV1#prefix}.
	Prefix *string `field:"required" json:"prefix" yaml:"prefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/css_cluster_v1#start_time CssClusterV1#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
}

