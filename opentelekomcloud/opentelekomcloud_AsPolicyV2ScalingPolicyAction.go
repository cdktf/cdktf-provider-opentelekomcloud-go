// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type AsPolicyV2ScalingPolicyAction struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#limits AsPolicyV2#limits}.
	Limits *float64 `field:"optional" json:"limits" yaml:"limits"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#operation AsPolicyV2#operation}.
	Operation *string `field:"optional" json:"operation" yaml:"operation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#percentage AsPolicyV2#percentage}.
	Percentage *float64 `field:"optional" json:"percentage" yaml:"percentage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/as_policy_v2#size AsPolicyV2#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
}
