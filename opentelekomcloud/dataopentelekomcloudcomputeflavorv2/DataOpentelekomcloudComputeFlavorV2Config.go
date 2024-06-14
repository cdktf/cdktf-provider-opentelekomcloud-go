// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudcomputeflavorv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudComputeFlavorV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/compute_flavor_v2#availability_zone DataOpentelekomcloudComputeFlavorV2#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/compute_flavor_v2#disk DataOpentelekomcloudComputeFlavorV2#disk}.
	Disk *float64 `field:"optional" json:"disk" yaml:"disk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/compute_flavor_v2#flavor_id DataOpentelekomcloudComputeFlavorV2#flavor_id}.
	FlavorId *string `field:"optional" json:"flavorId" yaml:"flavorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/compute_flavor_v2#id DataOpentelekomcloudComputeFlavorV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/compute_flavor_v2#min_disk DataOpentelekomcloudComputeFlavorV2#min_disk}.
	MinDisk *float64 `field:"optional" json:"minDisk" yaml:"minDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/compute_flavor_v2#min_ram DataOpentelekomcloudComputeFlavorV2#min_ram}.
	MinRam *float64 `field:"optional" json:"minRam" yaml:"minRam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/compute_flavor_v2#name DataOpentelekomcloudComputeFlavorV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/compute_flavor_v2#ram DataOpentelekomcloudComputeFlavorV2#ram}.
	Ram *float64 `field:"optional" json:"ram" yaml:"ram"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/compute_flavor_v2#region DataOpentelekomcloudComputeFlavorV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/compute_flavor_v2#resource_type DataOpentelekomcloudComputeFlavorV2#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/compute_flavor_v2#rx_tx_factor DataOpentelekomcloudComputeFlavorV2#rx_tx_factor}.
	RxTxFactor *float64 `field:"optional" json:"rxTxFactor" yaml:"rxTxFactor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/compute_flavor_v2#swap DataOpentelekomcloudComputeFlavorV2#swap}.
	Swap *float64 `field:"optional" json:"swap" yaml:"swap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/compute_flavor_v2#vcpus DataOpentelekomcloudComputeFlavorV2#vcpus}.
	Vcpus *float64 `field:"optional" json:"vcpus" yaml:"vcpus"`
}

