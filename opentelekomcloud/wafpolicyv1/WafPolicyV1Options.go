// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafpolicyv1


type WafPolicyV1Options struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#antitamper WafPolicyV1#antitamper}.
	Antitamper interface{} `field:"optional" json:"antitamper" yaml:"antitamper"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#cc WafPolicyV1#cc}.
	Cc interface{} `field:"optional" json:"cc" yaml:"cc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#common WafPolicyV1#common}.
	Common interface{} `field:"optional" json:"common" yaml:"common"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#crawler WafPolicyV1#crawler}.
	Crawler interface{} `field:"optional" json:"crawler" yaml:"crawler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#crawler_engine WafPolicyV1#crawler_engine}.
	CrawlerEngine interface{} `field:"optional" json:"crawlerEngine" yaml:"crawlerEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#crawler_other WafPolicyV1#crawler_other}.
	CrawlerOther interface{} `field:"optional" json:"crawlerOther" yaml:"crawlerOther"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#crawler_scanner WafPolicyV1#crawler_scanner}.
	CrawlerScanner interface{} `field:"optional" json:"crawlerScanner" yaml:"crawlerScanner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#crawler_script WafPolicyV1#crawler_script}.
	CrawlerScript interface{} `field:"optional" json:"crawlerScript" yaml:"crawlerScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#custom WafPolicyV1#custom}.
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#ignore WafPolicyV1#ignore}.
	Ignore interface{} `field:"optional" json:"ignore" yaml:"ignore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#privacy WafPolicyV1#privacy}.
	Privacy interface{} `field:"optional" json:"privacy" yaml:"privacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#webattack WafPolicyV1#webattack}.
	Webattack interface{} `field:"optional" json:"webattack" yaml:"webattack"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#webshell WafPolicyV1#webshell}.
	Webshell interface{} `field:"optional" json:"webshell" yaml:"webshell"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/waf_policy_v1#whiteblackip WafPolicyV1#whiteblackip}.
	Whiteblackip interface{} `field:"optional" json:"whiteblackip" yaml:"whiteblackip"`
}

