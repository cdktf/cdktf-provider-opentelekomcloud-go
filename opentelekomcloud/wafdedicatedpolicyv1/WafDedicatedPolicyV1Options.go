// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package wafdedicatedpolicyv1


type WafDedicatedPolicyV1Options struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#anti_crawler WafDedicatedPolicyV1#anti_crawler}.
	AntiCrawler interface{} `field:"optional" json:"antiCrawler" yaml:"antiCrawler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#anti_leakage WafDedicatedPolicyV1#anti_leakage}.
	AntiLeakage interface{} `field:"optional" json:"antiLeakage" yaml:"antiLeakage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#anti_tamper WafDedicatedPolicyV1#anti_tamper}.
	AntiTamper interface{} `field:"optional" json:"antiTamper" yaml:"antiTamper"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#blacklist WafDedicatedPolicyV1#blacklist}.
	Blacklist interface{} `field:"optional" json:"blacklist" yaml:"blacklist"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#cc WafDedicatedPolicyV1#cc}.
	Cc interface{} `field:"optional" json:"cc" yaml:"cc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#common WafDedicatedPolicyV1#common}.
	Common interface{} `field:"optional" json:"common" yaml:"common"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#crawler WafDedicatedPolicyV1#crawler}.
	Crawler interface{} `field:"optional" json:"crawler" yaml:"crawler"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#crawler_engine WafDedicatedPolicyV1#crawler_engine}.
	CrawlerEngine interface{} `field:"optional" json:"crawlerEngine" yaml:"crawlerEngine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#crawler_other WafDedicatedPolicyV1#crawler_other}.
	CrawlerOther interface{} `field:"optional" json:"crawlerOther" yaml:"crawlerOther"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#crawler_scanner WafDedicatedPolicyV1#crawler_scanner}.
	CrawlerScanner interface{} `field:"optional" json:"crawlerScanner" yaml:"crawlerScanner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#crawler_script WafDedicatedPolicyV1#crawler_script}.
	CrawlerScript interface{} `field:"optional" json:"crawlerScript" yaml:"crawlerScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#custom WafDedicatedPolicyV1#custom}.
	Custom interface{} `field:"optional" json:"custom" yaml:"custom"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#followed_action WafDedicatedPolicyV1#followed_action}.
	FollowedAction interface{} `field:"optional" json:"followedAction" yaml:"followedAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#geolocation_access_control WafDedicatedPolicyV1#geolocation_access_control}.
	GeolocationAccessControl interface{} `field:"optional" json:"geolocationAccessControl" yaml:"geolocationAccessControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#ignore WafDedicatedPolicyV1#ignore}.
	Ignore interface{} `field:"optional" json:"ignore" yaml:"ignore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#privacy WafDedicatedPolicyV1#privacy}.
	Privacy interface{} `field:"optional" json:"privacy" yaml:"privacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#web_attack WafDedicatedPolicyV1#web_attack}.
	WebAttack interface{} `field:"optional" json:"webAttack" yaml:"webAttack"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.19/docs/resources/waf_dedicated_policy_v1#web_shell WafDedicatedPolicyV1#web_shell}.
	WebShell interface{} `field:"optional" json:"webShell" yaml:"webShell"`
}

