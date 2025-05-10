// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudcomputebmsflavorsv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudComputeBmsFlavorsV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/data-sources/compute_bms_flavors_v2#id DataOpentelekomcloudComputeBmsFlavorsV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/data-sources/compute_bms_flavors_v2#min_disk DataOpentelekomcloudComputeBmsFlavorsV2#min_disk}.
	MinDisk *float64 `field:"optional" json:"minDisk" yaml:"minDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/data-sources/compute_bms_flavors_v2#min_ram DataOpentelekomcloudComputeBmsFlavorsV2#min_ram}.
	MinRam *float64 `field:"optional" json:"minRam" yaml:"minRam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/data-sources/compute_bms_flavors_v2#name DataOpentelekomcloudComputeBmsFlavorsV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/data-sources/compute_bms_flavors_v2#region DataOpentelekomcloudComputeBmsFlavorsV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/data-sources/compute_bms_flavors_v2#sort_dir DataOpentelekomcloudComputeBmsFlavorsV2#sort_dir}.
	SortDir *string `field:"optional" json:"sortDir" yaml:"sortDir"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.38/docs/data-sources/compute_bms_flavors_v2#sort_key DataOpentelekomcloudComputeBmsFlavorsV2#sort_key}.
	SortKey *string `field:"optional" json:"sortKey" yaml:"sortKey"`
}

