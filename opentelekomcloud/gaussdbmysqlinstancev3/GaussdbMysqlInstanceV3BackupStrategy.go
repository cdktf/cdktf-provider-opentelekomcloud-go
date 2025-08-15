// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gaussdbmysqlinstancev3


type GaussdbMysqlInstanceV3BackupStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/gaussdb_mysql_instance_v3#start_time GaussdbMysqlInstanceV3#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.45/docs/resources/gaussdb_mysql_instance_v3#keep_days GaussdbMysqlInstanceV3#keep_days}.
	KeepDays *string `field:"optional" json:"keepDays" yaml:"keepDays"`
}

