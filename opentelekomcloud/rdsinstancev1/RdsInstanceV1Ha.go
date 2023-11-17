// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsinstancev1


type RdsInstanceV1Ha struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.12/docs/resources/rds_instance_v1#enable RdsInstanceV1#enable}.
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.12/docs/resources/rds_instance_v1#replicationmode RdsInstanceV1#replicationmode}.
	Replicationmode *string `field:"optional" json:"replicationmode" yaml:"replicationmode"`
}

