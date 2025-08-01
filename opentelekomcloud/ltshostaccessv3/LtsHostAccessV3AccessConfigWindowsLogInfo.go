// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltshostaccessv3


type LtsHostAccessV3AccessConfigWindowsLogInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lts_host_access_v3#categories LtsHostAccessV3#categories}.
	Categories *[]*string `field:"required" json:"categories" yaml:"categories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lts_host_access_v3#event_level LtsHostAccessV3#event_level}.
	EventLevel *[]*string `field:"required" json:"eventLevel" yaml:"eventLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lts_host_access_v3#time_offset LtsHostAccessV3#time_offset}.
	TimeOffset *float64 `field:"required" json:"timeOffset" yaml:"timeOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.44/docs/resources/lts_host_access_v3#time_offset_unit LtsHostAccessV3#time_offset_unit}.
	TimeOffsetUnit *string `field:"required" json:"timeOffsetUnit" yaml:"timeOffsetUnit"`
}

