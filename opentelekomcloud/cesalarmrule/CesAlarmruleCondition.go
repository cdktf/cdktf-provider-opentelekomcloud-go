// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cesalarmrule


type CesAlarmruleCondition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/ces_alarmrule#comparison_operator CesAlarmrule#comparison_operator}.
	ComparisonOperator *string `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/ces_alarmrule#count CesAlarmrule#count}.
	Count *float64 `field:"required" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/ces_alarmrule#filter CesAlarmrule#filter}.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/ces_alarmrule#period CesAlarmrule#period}.
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/ces_alarmrule#value CesAlarmrule#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/ces_alarmrule#alarm_frequency CesAlarmrule#alarm_frequency}.
	AlarmFrequency *float64 `field:"optional" json:"alarmFrequency" yaml:"alarmFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.30/docs/resources/ces_alarmrule#unit CesAlarmrule#unit}.
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
}

