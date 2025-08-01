// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package computevolumeattachv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ComputeVolumeAttachV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/compute_volume_attach_v2#instance_id ComputeVolumeAttachV2#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/compute_volume_attach_v2#volume_id ComputeVolumeAttachV2#volume_id}.
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/compute_volume_attach_v2#device ComputeVolumeAttachV2#device}.
	Device *string `field:"optional" json:"device" yaml:"device"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/compute_volume_attach_v2#id ComputeVolumeAttachV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/compute_volume_attach_v2#region ComputeVolumeAttachV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/compute_volume_attach_v2#timeouts ComputeVolumeAttachV2#timeouts}
	Timeouts *ComputeVolumeAttachV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

