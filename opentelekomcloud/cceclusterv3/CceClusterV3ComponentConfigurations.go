// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cceclusterv3


type CceClusterV3ComponentConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_cluster_v3#name CceClusterV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_cluster_v3#configurations CceClusterV3#configurations}
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
}

