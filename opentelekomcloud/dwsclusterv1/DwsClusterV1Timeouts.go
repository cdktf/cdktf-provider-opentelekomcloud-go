// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dwsclusterv1


type DwsClusterV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/dws_cluster_v1#create DwsClusterV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/dws_cluster_v1#delete DwsClusterV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/dws_cluster_v1#update DwsClusterV1#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

