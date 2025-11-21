// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudcesmultiplemetricdatav1


type DataOpentelekomcloudCesMultipleMetricDataV1Metrics struct {
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/data-sources/ces_multiple_metric_data_v1#dimensions DataOpentelekomcloudCesMultipleMetricDataV1#dimensions}
	Dimensions interface{} `field:"required" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/data-sources/ces_multiple_metric_data_v1#metric_name DataOpentelekomcloudCesMultipleMetricDataV1#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/data-sources/ces_multiple_metric_data_v1#namespace DataOpentelekomcloudCesMultipleMetricDataV1#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

