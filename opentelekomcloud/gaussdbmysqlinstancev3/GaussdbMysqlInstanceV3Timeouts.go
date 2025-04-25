// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package gaussdbmysqlinstancev3


type GaussdbMysqlInstanceV3Timeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/gaussdb_mysql_instance_v3#create GaussdbMysqlInstanceV3#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/gaussdb_mysql_instance_v3#delete GaussdbMysqlInstanceV3#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.37/docs/resources/gaussdb_mysql_instance_v3#update GaussdbMysqlInstanceV3#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

