// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sfsfilesystemv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SfsFileSystemV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#size SfsFileSystemV2#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#access_level SfsFileSystemV2#access_level}.
	AccessLevel *string `field:"optional" json:"accessLevel" yaml:"accessLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#access_to SfsFileSystemV2#access_to}.
	AccessTo *string `field:"optional" json:"accessTo" yaml:"accessTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#access_type SfsFileSystemV2#access_type}.
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#availability_zone SfsFileSystemV2#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#description SfsFileSystemV2#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#id SfsFileSystemV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#is_public SfsFileSystemV2#is_public}.
	IsPublic interface{} `field:"optional" json:"isPublic" yaml:"isPublic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#metadata SfsFileSystemV2#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#name SfsFileSystemV2#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#region SfsFileSystemV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#share_proto SfsFileSystemV2#share_proto}.
	ShareProto *string `field:"optional" json:"shareProto" yaml:"shareProto"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#tags SfsFileSystemV2#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/sfs_file_system_v2#timeouts SfsFileSystemV2#timeouts}
	Timeouts *SfsFileSystemV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

