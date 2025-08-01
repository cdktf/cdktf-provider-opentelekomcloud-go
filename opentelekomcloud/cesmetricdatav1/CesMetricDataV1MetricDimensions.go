// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cesmetricdatav1


type CesMetricDataV1MetricDimensions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/ces_metric_data_v1#name CesMetricDataV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/ces_metric_data_v1#value CesMetricDataV1#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

