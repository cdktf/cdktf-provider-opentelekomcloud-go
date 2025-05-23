// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mrsclusterv1


type MrsClusterV1BootstrapScripts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/mrs_cluster_v1#fail_action MrsClusterV1#fail_action}.
	FailAction *string `field:"required" json:"failAction" yaml:"failAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/mrs_cluster_v1#name MrsClusterV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/mrs_cluster_v1#nodes MrsClusterV1#nodes}.
	Nodes *[]*string `field:"required" json:"nodes" yaml:"nodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/mrs_cluster_v1#uri MrsClusterV1#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/mrs_cluster_v1#active_master MrsClusterV1#active_master}.
	ActiveMaster interface{} `field:"optional" json:"activeMaster" yaml:"activeMaster"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/mrs_cluster_v1#before_component_start MrsClusterV1#before_component_start}.
	BeforeComponentStart interface{} `field:"optional" json:"beforeComponentStart" yaml:"beforeComponentStart"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/mrs_cluster_v1#parameters MrsClusterV1#parameters}.
	Parameters *string `field:"optional" json:"parameters" yaml:"parameters"`
}

