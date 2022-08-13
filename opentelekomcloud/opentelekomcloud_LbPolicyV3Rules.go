// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type LbPolicyV3Rules struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_policy_v3#compare_type LbPolicyV3#compare_type}.
	CompareType *string `field:"required" json:"compareType" yaml:"compareType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_policy_v3#type LbPolicyV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_policy_v3#value LbPolicyV3#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

