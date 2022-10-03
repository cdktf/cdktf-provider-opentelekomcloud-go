package aspolicyv1


type AsPolicyV1ScalingPolicyAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v1#instance_number AsPolicyV1#instance_number}.
	InstanceNumber *float64 `field:"optional" json:"instanceNumber" yaml:"instanceNumber"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v1#operation AsPolicyV1#operation}.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
}

