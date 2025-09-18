// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cesmetricdatav1


type CesMetricDataV1Metric struct {
	// dimensions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/ces_metric_data_v1#dimensions CesMetricDataV1#dimensions}
	Dimensions interface{} `field:"required" json:"dimensions" yaml:"dimensions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/ces_metric_data_v1#metric_name CesMetricDataV1#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/ces_metric_data_v1#namespace CesMetricDataV1#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
}

