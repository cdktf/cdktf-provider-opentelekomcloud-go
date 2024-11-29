// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package aspolicyv2


type AsPolicyV2ScheduledPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/as_policy_v2#launch_time AsPolicyV2#launch_time}.
	LaunchTime *string `field:"required" json:"launchTime" yaml:"launchTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/as_policy_v2#end_time AsPolicyV2#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/as_policy_v2#recurrence_type AsPolicyV2#recurrence_type}.
	RecurrenceType *string `field:"optional" json:"recurrenceType" yaml:"recurrenceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/as_policy_v2#recurrence_value AsPolicyV2#recurrence_value}.
	RecurrenceValue *string `field:"optional" json:"recurrenceValue" yaml:"recurrenceValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/as_policy_v2#start_time AsPolicyV2#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

