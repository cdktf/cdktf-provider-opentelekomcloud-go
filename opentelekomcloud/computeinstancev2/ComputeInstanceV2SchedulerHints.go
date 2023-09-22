// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computeinstancev2


type ComputeInstanceV2SchedulerHints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/compute_instance_v2#build_near_host_ip ComputeInstanceV2#build_near_host_ip}.
	BuildNearHostIp *string `field:"optional" json:"buildNearHostIp" yaml:"buildNearHostIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/compute_instance_v2#deh_id ComputeInstanceV2#deh_id}.
	DehId *string `field:"optional" json:"dehId" yaml:"dehId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/compute_instance_v2#different_host ComputeInstanceV2#different_host}.
	DifferentHost *[]*string `field:"optional" json:"differentHost" yaml:"differentHost"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/compute_instance_v2#group ComputeInstanceV2#group}.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/compute_instance_v2#query ComputeInstanceV2#query}.
	Query *[]*string `field:"optional" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/compute_instance_v2#same_host ComputeInstanceV2#same_host}.
	SameHost *[]*string `field:"optional" json:"sameHost" yaml:"sameHost"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/compute_instance_v2#target_cell ComputeInstanceV2#target_cell}.
	TargetCell *string `field:"optional" json:"targetCell" yaml:"targetCell"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/compute_instance_v2#tenancy ComputeInstanceV2#tenancy}.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

