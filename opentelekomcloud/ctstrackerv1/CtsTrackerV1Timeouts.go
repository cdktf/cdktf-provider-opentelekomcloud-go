// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ctstrackerv1


type CtsTrackerV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/cts_tracker_v1#create CtsTrackerV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/cts_tracker_v1#delete CtsTrackerV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

