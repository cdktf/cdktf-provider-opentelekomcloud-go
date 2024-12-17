// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sfsfilesystemv2


type SfsFileSystemV2Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/sfs_file_system_v2#create SfsFileSystemV2#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/sfs_file_system_v2#delete SfsFileSystemV2#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

