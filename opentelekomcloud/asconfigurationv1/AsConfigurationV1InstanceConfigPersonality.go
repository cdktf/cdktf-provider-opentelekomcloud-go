// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package asconfigurationv1


type AsConfigurationV1InstanceConfigPersonality struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/as_configuration_v1#content AsConfigurationV1#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.41/docs/resources/as_configuration_v1#path AsConfigurationV1#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
}

