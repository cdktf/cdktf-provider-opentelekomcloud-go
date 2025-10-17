// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudcesmetricdatav1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudCesMetricDataV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/ces_metric_data_v1#dim0 DataOpentelekomcloudCesMetricDataV1#dim0}.
	Dim0 *string `field:"required" json:"dim0" yaml:"dim0"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/ces_metric_data_v1#filter DataOpentelekomcloudCesMetricDataV1#filter}.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/ces_metric_data_v1#from DataOpentelekomcloudCesMetricDataV1#from}.
	From *string `field:"required" json:"from" yaml:"from"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/ces_metric_data_v1#metric_name DataOpentelekomcloudCesMetricDataV1#metric_name}.
	MetricName *string `field:"required" json:"metricName" yaml:"metricName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/ces_metric_data_v1#namespace DataOpentelekomcloudCesMetricDataV1#namespace}.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/ces_metric_data_v1#period DataOpentelekomcloudCesMetricDataV1#period}.
	Period *float64 `field:"required" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/ces_metric_data_v1#to DataOpentelekomcloudCesMetricDataV1#to}.
	To *string `field:"required" json:"to" yaml:"to"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/ces_metric_data_v1#dim1 DataOpentelekomcloudCesMetricDataV1#dim1}.
	Dim1 *string `field:"optional" json:"dim1" yaml:"dim1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/data-sources/ces_metric_data_v1#id DataOpentelekomcloudCesMetricDataV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

