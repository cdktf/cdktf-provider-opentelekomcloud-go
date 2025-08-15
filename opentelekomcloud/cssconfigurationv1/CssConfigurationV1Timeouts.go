// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cssconfigurationv1


type CssConfigurationV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/css_configuration_v1#create CssConfigurationV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/css_configuration_v1#delete CssConfigurationV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

