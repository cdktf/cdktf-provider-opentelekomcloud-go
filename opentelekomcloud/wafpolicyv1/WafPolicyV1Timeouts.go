package wafpolicyv1


type WafPolicyV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.5/docs/resources/waf_policy_v1#create WafPolicyV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.5/docs/resources/waf_policy_v1#delete WafPolicyV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

