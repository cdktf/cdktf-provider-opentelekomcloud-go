// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwapiv2


type ApigwApiV2HttpPolicyConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_api_v2#value ApigwApiV2#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_api_v2#origin ApigwApiV2#origin}.
	Origin *string `field:"optional" json:"origin" yaml:"origin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_api_v2#param_name ApigwApiV2#param_name}.
	ParamName *string `field:"optional" json:"paramName" yaml:"paramName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.39/docs/resources/apigw_api_v2#type ApigwApiV2#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

