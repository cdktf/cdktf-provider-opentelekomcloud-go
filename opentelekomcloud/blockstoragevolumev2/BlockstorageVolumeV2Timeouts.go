// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package blockstoragevolumev2


type BlockstorageVolumeV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/blockstorage_volume_v2#create BlockstorageVolumeV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.13/docs/resources/blockstorage_volume_v2#delete BlockstorageVolumeV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

