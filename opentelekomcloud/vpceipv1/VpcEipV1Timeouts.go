// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpceipv1


type VpcEipV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/resources/vpc_eip_v1#create VpcEipV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.11/docs/resources/vpc_eip_v1#delete VpcEipV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

