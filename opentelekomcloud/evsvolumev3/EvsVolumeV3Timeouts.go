// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package evsvolumev3


type EvsVolumeV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/evs_volume_v3#create EvsVolumeV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.23/docs/resources/evs_volume_v3#delete EvsVolumeV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

