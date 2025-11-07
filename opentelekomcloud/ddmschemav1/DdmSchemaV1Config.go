// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ddmschemav1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DdmSchemaV1Config struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/ddm_schema_v1#instance_id DdmSchemaV1#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/ddm_schema_v1#name DdmSchemaV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// rds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/ddm_schema_v1#rds DdmSchemaV1#rds}
	Rds interface{} `field:"required" json:"rds" yaml:"rds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/ddm_schema_v1#shard_mode DdmSchemaV1#shard_mode}.
	ShardMode *string `field:"required" json:"shardMode" yaml:"shardMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/ddm_schema_v1#shard_number DdmSchemaV1#shard_number}.
	ShardNumber *float64 `field:"required" json:"shardNumber" yaml:"shardNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/ddm_schema_v1#id DdmSchemaV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/ddm_schema_v1#purge_rds_on_delete DdmSchemaV1#purge_rds_on_delete}.
	PurgeRdsOnDelete interface{} `field:"optional" json:"purgeRdsOnDelete" yaml:"purgeRdsOnDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/ddm_schema_v1#shard_unit DdmSchemaV1#shard_unit}.
	ShardUnit *float64 `field:"optional" json:"shardUnit" yaml:"shardUnit"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/ddm_schema_v1#timeouts DdmSchemaV1#timeouts}
	Timeouts *DdmSchemaV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

