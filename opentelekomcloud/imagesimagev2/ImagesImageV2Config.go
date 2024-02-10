// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagesimagev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagesImageV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#container_format ImagesImageV2#container_format}.
	ContainerFormat *string `field:"required" json:"containerFormat" yaml:"containerFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#disk_format ImagesImageV2#disk_format}.
	DiskFormat *string `field:"required" json:"diskFormat" yaml:"diskFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#name ImagesImageV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#id ImagesImageV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#image_cache_path ImagesImageV2#image_cache_path}.
	ImageCachePath *string `field:"optional" json:"imageCachePath" yaml:"imageCachePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#image_source_url ImagesImageV2#image_source_url}.
	ImageSourceUrl *string `field:"optional" json:"imageSourceUrl" yaml:"imageSourceUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#local_file_path ImagesImageV2#local_file_path}.
	LocalFilePath *string `field:"optional" json:"localFilePath" yaml:"localFilePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#min_disk_gb ImagesImageV2#min_disk_gb}.
	MinDiskGb *float64 `field:"optional" json:"minDiskGb" yaml:"minDiskGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#min_ram_mb ImagesImageV2#min_ram_mb}.
	MinRamMb *float64 `field:"optional" json:"minRamMb" yaml:"minRamMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#protected ImagesImageV2#protected}.
	Protected interface{} `field:"optional" json:"protected" yaml:"protected"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#region ImagesImageV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#tags ImagesImageV2#tags}.
	Tags *[]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#timeouts ImagesImageV2#timeouts}
	Timeouts *ImagesImageV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.1/docs/resources/images_image_v2#visibility ImagesImageV2#visibility}.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
}

