// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package antiddosv1


type AntiddosV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/antiddos_v1#create AntiddosV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/antiddos_v1#delete AntiddosV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/antiddos_v1#update AntiddosV1#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

