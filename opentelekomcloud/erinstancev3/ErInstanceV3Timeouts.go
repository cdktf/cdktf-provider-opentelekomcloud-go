// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package erinstancev3


type ErInstanceV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/er_instance_v3#create ErInstanceV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/er_instance_v3#delete ErInstanceV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/resources/er_instance_v3#update ErInstanceV3#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

