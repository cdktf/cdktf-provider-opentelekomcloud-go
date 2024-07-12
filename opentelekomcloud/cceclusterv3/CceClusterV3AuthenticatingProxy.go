// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cceclusterv3


type CceClusterV3AuthenticatingProxy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/cce_cluster_v3#ca CceClusterV3#ca}.
	Ca *string `field:"required" json:"ca" yaml:"ca"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/cce_cluster_v3#cert CceClusterV3#cert}.
	Cert *string `field:"required" json:"cert" yaml:"cert"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.14/docs/resources/cce_cluster_v3#private_key CceClusterV3#private_key}.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
}

