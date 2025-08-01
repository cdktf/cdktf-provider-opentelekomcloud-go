// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ddsltslogv3


type DdsLtsLogV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/dds_lts_log_v3#create DdsLtsLogV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/dds_lts_log_v3#delete DdsLtsLogV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

