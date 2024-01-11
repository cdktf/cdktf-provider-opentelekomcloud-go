// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gaussdbmysqlinstancev3


type GaussdbMysqlInstanceV3Datastore struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/gaussdb_mysql_instance_v3#engine GaussdbMysqlInstanceV3#engine}.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/gaussdb_mysql_instance_v3#version GaussdbMysqlInstanceV3#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
}

