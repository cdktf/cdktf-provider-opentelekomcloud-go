// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aspolicyv1


type AsPolicyV1ScalingPolicyAction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/as_policy_v1#instance_number AsPolicyV1#instance_number}.
	InstanceNumber *float64 `field:"optional" json:"instanceNumber" yaml:"instanceNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/resources/as_policy_v1#operation AsPolicyV1#operation}.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
}

