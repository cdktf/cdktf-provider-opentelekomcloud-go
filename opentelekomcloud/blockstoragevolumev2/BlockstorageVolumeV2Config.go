// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package blockstoragevolumev2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BlockstorageVolumeV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#size BlockstorageVolumeV2#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#availability_zone BlockstorageVolumeV2#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#cascade BlockstorageVolumeV2#cascade}.
	Cascade interface{} `field:"optional" json:"cascade" yaml:"cascade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#consistency_group_id BlockstorageVolumeV2#consistency_group_id}.
	ConsistencyGroupId *string `field:"optional" json:"consistencyGroupId" yaml:"consistencyGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#description BlockstorageVolumeV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#device_type BlockstorageVolumeV2#device_type}.
	DeviceType *string `field:"optional" json:"deviceType" yaml:"deviceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#id BlockstorageVolumeV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#image_id BlockstorageVolumeV2#image_id}.
	ImageId *string `field:"optional" json:"imageId" yaml:"imageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#metadata BlockstorageVolumeV2#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#name BlockstorageVolumeV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#region BlockstorageVolumeV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#snapshot_id BlockstorageVolumeV2#snapshot_id}.
	SnapshotId *string `field:"optional" json:"snapshotId" yaml:"snapshotId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#source_replica BlockstorageVolumeV2#source_replica}.
	SourceReplica *string `field:"optional" json:"sourceReplica" yaml:"sourceReplica"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#source_vol_id BlockstorageVolumeV2#source_vol_id}.
	SourceVolId *string `field:"optional" json:"sourceVolId" yaml:"sourceVolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#tags BlockstorageVolumeV2#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#timeouts BlockstorageVolumeV2#timeouts}
	Timeouts *BlockstorageVolumeV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/blockstorage_volume_v2#volume_type BlockstorageVolumeV2#volume_type}.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

