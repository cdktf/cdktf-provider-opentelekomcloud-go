// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwcustomauthorizerv2


type ApigwCustomAuthorizerV2Identity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.20/docs/resources/apigw_custom_authorizer_v2#location ApigwCustomAuthorizerV2#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.20/docs/resources/apigw_custom_authorizer_v2#name ApigwCustomAuthorizerV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.20/docs/resources/apigw_custom_authorizer_v2#validation ApigwCustomAuthorizerV2#validation}.
	Validation *string `field:"optional" json:"validation" yaml:"validation"`
}

