// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cssclusterv1


type CssClusterV1NodeConfigNetworkInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/css_cluster_v1#network_id CssClusterV1#network_id}.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/css_cluster_v1#security_group_id CssClusterV1#security_group_id}.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/css_cluster_v1#vpc_id CssClusterV1#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
}

