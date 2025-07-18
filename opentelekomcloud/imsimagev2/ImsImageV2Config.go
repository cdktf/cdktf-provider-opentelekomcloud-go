// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imsimagev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImsImageV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#name ImsImageV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#cmk_id ImsImageV2#cmk_id}.
	CmkId *string `field:"optional" json:"cmkId" yaml:"cmkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#description ImsImageV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#hw_firmware_type ImsImageV2#hw_firmware_type}.
	HwFirmwareType *string `field:"optional" json:"hwFirmwareType" yaml:"hwFirmwareType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#id ImsImageV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#image_url ImsImageV2#image_url}.
	ImageUrl *string `field:"optional" json:"imageUrl" yaml:"imageUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#instance_id ImsImageV2#instance_id}.
	InstanceId *string `field:"optional" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#is_config ImsImageV2#is_config}.
	IsConfig interface{} `field:"optional" json:"isConfig" yaml:"isConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#max_ram ImsImageV2#max_ram}.
	MaxRam *float64 `field:"optional" json:"maxRam" yaml:"maxRam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#min_disk ImsImageV2#min_disk}.
	MinDisk *float64 `field:"optional" json:"minDisk" yaml:"minDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#min_ram ImsImageV2#min_ram}.
	MinRam *float64 `field:"optional" json:"minRam" yaml:"minRam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#os_version ImsImageV2#os_version}.
	OsVersion *string `field:"optional" json:"osVersion" yaml:"osVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#tags ImsImageV2#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#timeouts ImsImageV2#timeouts}
	Timeouts *ImsImageV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#type ImsImageV2#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/ims_image_v2#volume_id ImsImageV2#volume_id}.
	VolumeId *string `field:"optional" json:"volumeId" yaml:"volumeId"`
}

