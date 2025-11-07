// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudcesmultiplemetricdatav1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudCesMultipleMetricDataV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/data-sources/ces_multiple_metric_data_v1#filter DataOpentelekomcloudCesMultipleMetricDataV1#filter}.
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/data-sources/ces_multiple_metric_data_v1#from DataOpentelekomcloudCesMultipleMetricDataV1#from}.
	From *float64 `field:"required" json:"from" yaml:"from"`
	// metrics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/data-sources/ces_multiple_metric_data_v1#metrics DataOpentelekomcloudCesMultipleMetricDataV1#metrics}
	Metrics interface{} `field:"required" json:"metrics" yaml:"metrics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/data-sources/ces_multiple_metric_data_v1#period DataOpentelekomcloudCesMultipleMetricDataV1#period}.
	Period *string `field:"required" json:"period" yaml:"period"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/data-sources/ces_multiple_metric_data_v1#to DataOpentelekomcloudCesMultipleMetricDataV1#to}.
	To *float64 `field:"required" json:"to" yaml:"to"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/data-sources/ces_multiple_metric_data_v1#id DataOpentelekomcloudCesMultipleMetricDataV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

