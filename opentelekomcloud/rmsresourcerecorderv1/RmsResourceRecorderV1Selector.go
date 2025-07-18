// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rmsresourcerecorderv1


type RmsResourceRecorderV1Selector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/rms_resource_recorder_v1#all_supported RmsResourceRecorderV1#all_supported}.
	AllSupported interface{} `field:"required" json:"allSupported" yaml:"allSupported"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.43/docs/resources/rms_resource_recorder_v1#resource_types RmsResourceRecorderV1#resource_types}.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
}

