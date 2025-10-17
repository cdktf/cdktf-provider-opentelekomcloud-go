// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package taurusdbmysqlinstancev3


type TaurusdbMysqlInstanceV3Datastore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/taurusdb_mysql_instance_v3#engine TaurusdbMysqlInstanceV3#engine}.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/taurusdb_mysql_instance_v3#version TaurusdbMysqlInstanceV3#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

