// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cesmetricdatav1


type CesMetricDataV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/ces_metric_data_v1#create CesMetricDataV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/ces_metric_data_v1#delete CesMetricDataV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

