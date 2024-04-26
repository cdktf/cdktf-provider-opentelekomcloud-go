// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cesalarmrule


type CesAlarmruleMetric struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/ces_alarmrule#metric_name CesAlarmrule#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/ces_alarmrule#namespace CesAlarmrule#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/ces_alarmrule#dimensions CesAlarmrule#dimensions}
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
}

