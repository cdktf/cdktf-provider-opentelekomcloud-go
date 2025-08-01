// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imsimagev2


type ImsImageV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/ims_image_v2#create ImsImageV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/ims_image_v2#delete ImsImageV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

