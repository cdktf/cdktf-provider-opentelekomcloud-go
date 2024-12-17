// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigwthrottlingpolicyv2


type ApigwThrottlingPolicyV2AppThrottles struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/apigw_throttling_policy_v2#max_api_requests ApigwThrottlingPolicyV2#max_api_requests}.
	MaxApiRequests *float64 `field:"required" json:"maxApiRequests" yaml:"maxApiRequests"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/apigw_throttling_policy_v2#throttling_object_id ApigwThrottlingPolicyV2#throttling_object_id}.
	ThrottlingObjectId *string `field:"required" json:"throttlingObjectId" yaml:"throttlingObjectId"`
}

