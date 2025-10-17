// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mrsjobv1


type MrsJobV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/mrs_job_v1#create MrsJobV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/mrs_job_v1#delete MrsJobV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/mrs_job_v1#update MrsJobV1#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

