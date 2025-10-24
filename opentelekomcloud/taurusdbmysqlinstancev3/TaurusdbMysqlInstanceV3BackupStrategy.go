// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package taurusdbmysqlinstancev3


type TaurusdbMysqlInstanceV3BackupStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/taurusdb_mysql_instance_v3#start_time TaurusdbMysqlInstanceV3#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.51/docs/resources/taurusdb_mysql_instance_v3#keep_days TaurusdbMysqlInstanceV3#keep_days}.
	KeepDays *float64 `field:"optional" json:"keepDays" yaml:"keepDays"`
}

