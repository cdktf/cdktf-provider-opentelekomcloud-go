// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltshostaccessv3


type LtsHostAccessV3AccessConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/lts_host_access_v3#paths LtsHostAccessV3#paths}.
	Paths *[]*string `field:"required" json:"paths" yaml:"paths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/lts_host_access_v3#black_paths LtsHostAccessV3#black_paths}.
	BlackPaths *[]*string `field:"optional" json:"blackPaths" yaml:"blackPaths"`
	// multi_log_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/lts_host_access_v3#multi_log_format LtsHostAccessV3#multi_log_format}
	MultiLogFormat *LtsHostAccessV3AccessConfigMultiLogFormat `field:"optional" json:"multiLogFormat" yaml:"multiLogFormat"`
	// single_log_format block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/lts_host_access_v3#single_log_format LtsHostAccessV3#single_log_format}
	SingleLogFormat *LtsHostAccessV3AccessConfigSingleLogFormat `field:"optional" json:"singleLogFormat" yaml:"singleLogFormat"`
	// windows_log_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/lts_host_access_v3#windows_log_info LtsHostAccessV3#windows_log_info}
	WindowsLogInfo *LtsHostAccessV3AccessConfigWindowsLogInfo `field:"optional" json:"windowsLogInfo" yaml:"windowsLogInfo"`
}

