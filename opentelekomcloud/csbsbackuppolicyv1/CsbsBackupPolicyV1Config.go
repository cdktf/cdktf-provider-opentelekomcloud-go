// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package csbsbackuppolicyv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CsbsBackupPolicyV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/csbs_backup_policy_v1#name CsbsBackupPolicyV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/csbs_backup_policy_v1#resource CsbsBackupPolicyV1#resource}
	Resource interface{} `field:"required" json:"resource" yaml:"resource"`
	// scheduled_operation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/csbs_backup_policy_v1#scheduled_operation CsbsBackupPolicyV1#scheduled_operation}
	ScheduledOperation *CsbsBackupPolicyV1ScheduledOperation `field:"required" json:"scheduledOperation" yaml:"scheduledOperation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/csbs_backup_policy_v1#common CsbsBackupPolicyV1#common}.
	Common *map[string]*string `field:"optional" json:"common" yaml:"common"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/csbs_backup_policy_v1#description CsbsBackupPolicyV1#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/csbs_backup_policy_v1#id CsbsBackupPolicyV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/csbs_backup_policy_v1#provider_id CsbsBackupPolicyV1#provider_id}.
	ProviderId *string `field:"optional" json:"providerId" yaml:"providerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/csbs_backup_policy_v1#region CsbsBackupPolicyV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/csbs_backup_policy_v1#tags CsbsBackupPolicyV1#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.6/docs/resources/csbs_backup_policy_v1#timeouts CsbsBackupPolicyV1#timeouts}
	Timeouts *CsbsBackupPolicyV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

