// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltskeywordsalarmrulev2


type LtsKeywordsAlarmRuleV2KeywordsRequests struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_keywords_alarm_rule_v2#condition LtsKeywordsAlarmRuleV2#condition}.
	Condition *string `field:"required" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_keywords_alarm_rule_v2#keyword LtsKeywordsAlarmRuleV2#keyword}.
	Keyword *string `field:"required" json:"keyword" yaml:"keyword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_keywords_alarm_rule_v2#log_group_id LtsKeywordsAlarmRuleV2#log_group_id}.
	LogGroupId *string `field:"required" json:"logGroupId" yaml:"logGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_keywords_alarm_rule_v2#log_stream_id LtsKeywordsAlarmRuleV2#log_stream_id}.
	LogStreamId *string `field:"required" json:"logStreamId" yaml:"logStreamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_keywords_alarm_rule_v2#number LtsKeywordsAlarmRuleV2#number}.
	Number *float64 `field:"required" json:"number" yaml:"number"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_keywords_alarm_rule_v2#search_time_range LtsKeywordsAlarmRuleV2#search_time_range}.
	SearchTimeRange *float64 `field:"required" json:"searchTimeRange" yaml:"searchTimeRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/lts_keywords_alarm_rule_v2#search_time_range_unit LtsKeywordsAlarmRuleV2#search_time_range_unit}.
	SearchTimeRangeUnit *string `field:"required" json:"searchTimeRangeUnit" yaml:"searchTimeRangeUnit"`
}

