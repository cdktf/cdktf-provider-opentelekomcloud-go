package wafpreciseprotectionrulev1


type WafPreciseprotectionRuleV1Conditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/waf_preciseprotection_rule_v1#category WafPreciseprotectionRuleV1#category}.
	Category *string `field:"required" json:"category" yaml:"category"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/waf_preciseprotection_rule_v1#contents WafPreciseprotectionRuleV1#contents}.
	Contents *[]*string `field:"required" json:"contents" yaml:"contents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/waf_preciseprotection_rule_v1#logic WafPreciseprotectionRuleV1#logic}.
	Logic *string `field:"required" json:"logic" yaml:"logic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/waf_preciseprotection_rule_v1#index WafPreciseprotectionRuleV1#index}.
	Index *string `field:"optional" json:"index" yaml:"index"`
}

