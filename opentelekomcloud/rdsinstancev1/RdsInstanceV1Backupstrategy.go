// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsinstancev1


type RdsInstanceV1Backupstrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/rds_instance_v1#keepdays RdsInstanceV1#keepdays}.
	Keepdays *float64 `field:"optional" json:"keepdays" yaml:"keepdays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.27/docs/resources/rds_instance_v1#starttime RdsInstanceV1#starttime}.
	Starttime *string `field:"optional" json:"starttime" yaml:"starttime"`
}

