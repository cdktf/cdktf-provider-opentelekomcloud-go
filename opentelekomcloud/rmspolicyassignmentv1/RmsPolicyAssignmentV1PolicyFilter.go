// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rmspolicyassignmentv1


type RmsPolicyAssignmentV1PolicyFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/rms_policy_assignment_v1#region RmsPolicyAssignmentV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/rms_policy_assignment_v1#resource_id RmsPolicyAssignmentV1#resource_id}.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/rms_policy_assignment_v1#resource_provider RmsPolicyAssignmentV1#resource_provider}.
	ResourceProvider *string `field:"optional" json:"resourceProvider" yaml:"resourceProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/rms_policy_assignment_v1#resource_type RmsPolicyAssignmentV1#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/rms_policy_assignment_v1#tag_key RmsPolicyAssignmentV1#tag_key}.
	TagKey *string `field:"optional" json:"tagKey" yaml:"tagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/rms_policy_assignment_v1#tag_value RmsPolicyAssignmentV1#tag_value}.
	TagValue *string `field:"optional" json:"tagValue" yaml:"tagValue"`
}

