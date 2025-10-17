// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package drstaskv3


type DrsTaskV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/drs_task_v3#create DrsTaskV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/drs_task_v3#delete DrsTaskV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/drs_task_v3#update DrsTaskV3#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

