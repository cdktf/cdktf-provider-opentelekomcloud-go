// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ltscceaccessv3


type LtsCceAccessV3AccessConfigWindowsLogInfo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/lts_cce_access_v3#categories LtsCceAccessV3#categories}.
	Categories *[]*string `field:"required" json:"categories" yaml:"categories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/lts_cce_access_v3#event_level LtsCceAccessV3#event_level}.
	EventLevel *[]*string `field:"required" json:"eventLevel" yaml:"eventLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/lts_cce_access_v3#time_offset LtsCceAccessV3#time_offset}.
	TimeOffset *float64 `field:"required" json:"timeOffset" yaml:"timeOffset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.35/docs/resources/lts_cce_access_v3#time_offset_unit LtsCceAccessV3#time_offset_unit}.
	TimeOffsetUnit *string `field:"required" json:"timeOffsetUnit" yaml:"timeOffsetUnit"`
}

