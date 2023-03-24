package lbrulev3


type LbRuleV3Conditions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_rule_v3#value LbRuleV3#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/lb_rule_v3#key LbRuleV3#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

