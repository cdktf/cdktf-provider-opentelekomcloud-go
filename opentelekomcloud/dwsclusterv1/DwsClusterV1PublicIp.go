// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dwsclusterv1


type DwsClusterV1PublicIp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/dws_cluster_v1#eip_id DwsClusterV1#eip_id}.
	EipId *string `field:"optional" json:"eipId" yaml:"eipId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/dws_cluster_v1#public_bind_type DwsClusterV1#public_bind_type}.
	PublicBindType *string `field:"optional" json:"publicBindType" yaml:"publicBindType"`
}

