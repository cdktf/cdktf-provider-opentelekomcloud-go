// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dehhostv1


type DehHostV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/deh_host_v1#create DehHostV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/deh_host_v1#delete DehHostV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

