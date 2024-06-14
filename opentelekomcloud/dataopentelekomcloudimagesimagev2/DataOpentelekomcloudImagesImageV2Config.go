// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudimagesimagev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudImagesImageV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/images_image_v2#most_recent DataOpentelekomcloudImagesImageV2#most_recent}.
	MostRecent interface{} `field:"optional" json:"mostRecent" yaml:"mostRecent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/images_image_v2#name DataOpentelekomcloudImagesImageV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/images_image_v2#name_regex DataOpentelekomcloudImagesImageV2#name_regex}.
	NameRegex *string `field:"optional" json:"nameRegex" yaml:"nameRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/images_image_v2#owner DataOpentelekomcloudImagesImageV2#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/images_image_v2#size_max DataOpentelekomcloudImagesImageV2#size_max}.
	SizeMax *float64 `field:"optional" json:"sizeMax" yaml:"sizeMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/images_image_v2#size_min DataOpentelekomcloudImagesImageV2#size_min}.
	SizeMin *float64 `field:"optional" json:"sizeMin" yaml:"sizeMin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/images_image_v2#sort_direction DataOpentelekomcloudImagesImageV2#sort_direction}.
	SortDirection *string `field:"optional" json:"sortDirection" yaml:"sortDirection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/images_image_v2#sort_key DataOpentelekomcloudImagesImageV2#sort_key}.
	SortKey *string `field:"optional" json:"sortKey" yaml:"sortKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/images_image_v2#tag DataOpentelekomcloudImagesImageV2#tag}.
	Tag *string `field:"optional" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.12/docs/data-sources/images_image_v2#visibility DataOpentelekomcloudImagesImageV2#visibility}.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

