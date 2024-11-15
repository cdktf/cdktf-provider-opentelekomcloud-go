// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rtssoftwareconfigv1


type RtsSoftwareConfigV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/rts_software_config_v1#create RtsSoftwareConfigV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.25/docs/resources/rts_software_config_v1#delete RtsSoftwareConfigV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

