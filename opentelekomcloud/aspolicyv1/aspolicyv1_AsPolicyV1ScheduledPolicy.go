package aspolicyv1


type AsPolicyV1ScheduledPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v1#launch_time AsPolicyV1#launch_time}.
	LaunchTime *string `field:"required" json:"launchTime" yaml:"launchTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v1#end_time AsPolicyV1#end_time}.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v1#recurrence_type AsPolicyV1#recurrence_type}.
	RecurrenceType *string `field:"optional" json:"recurrenceType" yaml:"recurrenceType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v1#recurrence_value AsPolicyV1#recurrence_value}.
	RecurrenceValue *string `field:"optional" json:"recurrenceValue" yaml:"recurrenceValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v1#start_time AsPolicyV1#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
}

