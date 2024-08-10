// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpcsubnetv1


type VpcSubnetV1Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.16/docs/resources/vpc_subnet_v1#create VpcSubnetV1#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.16/docs/resources/vpc_subnet_v1#delete VpcSubnetV1#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

