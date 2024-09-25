// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cssclusterv1


type CssClusterV1NodeConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.20/docs/resources/css_cluster_v1#flavor CssClusterV1#flavor}.
	Flavor *string `field:"required" json:"flavor" yaml:"flavor"`
	// network_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.20/docs/resources/css_cluster_v1#network_info CssClusterV1#network_info}
	NetworkInfo *CssClusterV1NodeConfigNetworkInfo `field:"required" json:"networkInfo" yaml:"networkInfo"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.20/docs/resources/css_cluster_v1#volume CssClusterV1#volume}
	Volume *CssClusterV1NodeConfigVolume `field:"required" json:"volume" yaml:"volume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.20/docs/resources/css_cluster_v1#availability_zone CssClusterV1#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
}

