// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package drstaskv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DrsTaskV3Config struct {
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
	// destination_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#destination_db DrsTaskV3#destination_db}
	DestinationDb *DrsTaskV3DestinationDb `field:"required" json:"destinationDb" yaml:"destinationDb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#direction DrsTaskV3#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#engine_type DrsTaskV3#engine_type}.
	EngineType *string `field:"required" json:"engineType" yaml:"engineType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#name DrsTaskV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// source_db block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#source_db DrsTaskV3#source_db}
	SourceDb *DrsTaskV3SourceDb `field:"required" json:"sourceDb" yaml:"sourceDb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#type DrsTaskV3#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#description DrsTaskV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#destination_db_readonly DrsTaskV3#destination_db_readonly}.
	DestinationDbReadonly interface{} `field:"optional" json:"destinationDbReadonly" yaml:"destinationDbReadonly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#expired_days DrsTaskV3#expired_days}.
	ExpiredDays *float64 `field:"optional" json:"expiredDays" yaml:"expiredDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#force_destroy DrsTaskV3#force_destroy}.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#id DrsTaskV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// limit_speed block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#limit_speed DrsTaskV3#limit_speed}
	LimitSpeed interface{} `field:"optional" json:"limitSpeed" yaml:"limitSpeed"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#migrate_definer DrsTaskV3#migrate_definer}.
	MigrateDefiner interface{} `field:"optional" json:"migrateDefiner" yaml:"migrateDefiner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#migration_type DrsTaskV3#migration_type}.
	MigrationType *string `field:"optional" json:"migrationType" yaml:"migrationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#multi_write DrsTaskV3#multi_write}.
	MultiWrite interface{} `field:"optional" json:"multiWrite" yaml:"multiWrite"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#net_type DrsTaskV3#net_type}.
	NetType *string `field:"optional" json:"netType" yaml:"netType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#node_num DrsTaskV3#node_num}.
	NodeNum *float64 `field:"optional" json:"nodeNum" yaml:"nodeNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#start_time DrsTaskV3#start_time}.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#tags DrsTaskV3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.8/docs/resources/drs_task_v3#timeouts DrsTaskV3#timeouts}
	Timeouts *DrsTaskV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

