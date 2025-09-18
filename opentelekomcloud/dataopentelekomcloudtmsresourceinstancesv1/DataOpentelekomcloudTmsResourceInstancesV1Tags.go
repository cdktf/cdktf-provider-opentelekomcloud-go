// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudtmsresourceinstancesv1


type DataOpentelekomcloudTmsResourceInstancesV1Tags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/tms_resource_instances_v1#key DataOpentelekomcloudTmsResourceInstancesV1#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/data-sources/tms_resource_instances_v1#values DataOpentelekomcloudTmsResourceInstancesV1#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

