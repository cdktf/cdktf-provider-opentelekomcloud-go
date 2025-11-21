// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwresponsev2


type ApigwResponseV2Rule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_response_v2#body ApigwResponseV2#body}.
	Body *string `field:"required" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_response_v2#error_type ApigwResponseV2#error_type}.
	ErrorType *string `field:"required" json:"errorType" yaml:"errorType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/apigw_response_v2#status_code ApigwResponseV2#status_code}.
	StatusCode *float64 `field:"optional" json:"statusCode" yaml:"statusCode"`
}

