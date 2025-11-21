// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package geminitemplatev3


type GeminiTemplateV3Parameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/gemini_template_v3#name GeminiTemplateV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/gemini_template_v3#value GeminiTemplateV3#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

