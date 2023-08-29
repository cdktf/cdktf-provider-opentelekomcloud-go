// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vbsbackuppolicyv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VbsBackupPolicyV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vbs_backup_policy_v2#name VbsBackupPolicyV2#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vbs_backup_policy_v2#retain_first_backup VbsBackupPolicyV2#retain_first_backup}.
	RetainFirstBackup *string `field:"required" json:"retainFirstBackup" yaml:"retainFirstBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vbs_backup_policy_v2#start_time VbsBackupPolicyV2#start_time}.
	StartTime *string `field:"required" json:"startTime" yaml:"startTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vbs_backup_policy_v2#frequency VbsBackupPolicyV2#frequency}.
	Frequency *float64 `field:"optional" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vbs_backup_policy_v2#id VbsBackupPolicyV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vbs_backup_policy_v2#region VbsBackupPolicyV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vbs_backup_policy_v2#rentention_day VbsBackupPolicyV2#rentention_day}.
	RententionDay *float64 `field:"optional" json:"rententionDay" yaml:"rententionDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vbs_backup_policy_v2#rentention_num VbsBackupPolicyV2#rentention_num}.
	RententionNum *float64 `field:"optional" json:"rententionNum" yaml:"rententionNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vbs_backup_policy_v2#resources VbsBackupPolicyV2#resources}.
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vbs_backup_policy_v2#status VbsBackupPolicyV2#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vbs_backup_policy_v2#tags VbsBackupPolicyV2#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vbs_backup_policy_v2#timeouts VbsBackupPolicyV2#timeouts}
	Timeouts *VbsBackupPolicyV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/vbs_backup_policy_v2#week_frequency VbsBackupPolicyV2#week_frequency}.
	WeekFrequency *[]*string `field:"optional" json:"weekFrequency" yaml:"weekFrequency"`
}

