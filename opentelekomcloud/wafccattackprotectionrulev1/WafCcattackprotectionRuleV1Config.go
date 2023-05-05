package wafccattackprotectionrulev1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WafCcattackprotectionRuleV1Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#action_category WafCcattackprotectionRuleV1#action_category}.
	ActionCategory *string `field:"required" json:"actionCategory" yaml:"actionCategory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#limit_num WafCcattackprotectionRuleV1#limit_num}.
	LimitNum *float64 `field:"required" json:"limitNum" yaml:"limitNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#limit_period WafCcattackprotectionRuleV1#limit_period}.
	LimitPeriod *float64 `field:"required" json:"limitPeriod" yaml:"limitPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#policy_id WafCcattackprotectionRuleV1#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#tag_type WafCcattackprotectionRuleV1#tag_type}.
	TagType *string `field:"required" json:"tagType" yaml:"tagType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#url WafCcattackprotectionRuleV1#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#block_content WafCcattackprotectionRuleV1#block_content}.
	BlockContent *string `field:"optional" json:"blockContent" yaml:"blockContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#block_content_type WafCcattackprotectionRuleV1#block_content_type}.
	BlockContentType *string `field:"optional" json:"blockContentType" yaml:"blockContentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#id WafCcattackprotectionRuleV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#lock_time WafCcattackprotectionRuleV1#lock_time}.
	LockTime *float64 `field:"optional" json:"lockTime" yaml:"lockTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#tag_category WafCcattackprotectionRuleV1#tag_category}.
	TagCategory *string `field:"optional" json:"tagCategory" yaml:"tagCategory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#tag_contents WafCcattackprotectionRuleV1#tag_contents}.
	TagContents *[]*string `field:"optional" json:"tagContents" yaml:"tagContents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#tag_index WafCcattackprotectionRuleV1#tag_index}.
	TagIndex *string `field:"optional" json:"tagIndex" yaml:"tagIndex"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.3/docs/resources/waf_ccattackprotection_rule_v1#timeouts WafCcattackprotectionRuleV1#timeouts}
	Timeouts *WafCcattackprotectionRuleV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

