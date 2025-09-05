// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rmspolicyassignmentv1


type RmsPolicyAssignmentV1CustomPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/rms_policy_assignment_v1#auth_type RmsPolicyAssignmentV1#auth_type}.
	AuthType *string `field:"required" json:"authType" yaml:"authType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/rms_policy_assignment_v1#function_urn RmsPolicyAssignmentV1#function_urn}.
	FunctionUrn *string `field:"required" json:"functionUrn" yaml:"functionUrn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/rms_policy_assignment_v1#auth_value RmsPolicyAssignmentV1#auth_value}.
	AuthValue *map[string]*string `field:"optional" json:"authValue" yaml:"authValue"`
}

