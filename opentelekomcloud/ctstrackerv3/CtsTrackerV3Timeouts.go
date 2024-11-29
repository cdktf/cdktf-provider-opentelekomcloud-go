// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ctstrackerv3


type CtsTrackerV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cts_tracker_v3#create CtsTrackerV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cts_tracker_v3#delete CtsTrackerV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

