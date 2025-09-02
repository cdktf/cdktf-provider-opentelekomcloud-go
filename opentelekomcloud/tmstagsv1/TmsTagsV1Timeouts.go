// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tmstagsv1


type TmsTagsV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/tms_tags_v1#create TmsTagsV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.46/docs/resources/tms_tags_v1#delete TmsTagsV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

