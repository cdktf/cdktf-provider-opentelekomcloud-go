// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computevolumeattachv2


type ComputeVolumeAttachV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/compute_volume_attach_v2#create ComputeVolumeAttachV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/compute_volume_attach_v2#delete ComputeVolumeAttachV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

