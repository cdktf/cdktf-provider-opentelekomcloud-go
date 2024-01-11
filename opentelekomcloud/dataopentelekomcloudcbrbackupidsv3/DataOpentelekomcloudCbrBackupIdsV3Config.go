// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudcbrbackupidsv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudCbrBackupIdsV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/data-sources/cbr_backup_ids_v3#checkpoint_id DataOpentelekomcloudCbrBackupIdsV3#checkpoint_id}.
	CheckpointId *string `field:"optional" json:"checkpointId" yaml:"checkpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/data-sources/cbr_backup_ids_v3#id DataOpentelekomcloudCbrBackupIdsV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/data-sources/cbr_backup_ids_v3#image_type DataOpentelekomcloudCbrBackupIdsV3#image_type}.
	ImageType *string `field:"optional" json:"imageType" yaml:"imageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/data-sources/cbr_backup_ids_v3#name DataOpentelekomcloudCbrBackupIdsV3#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/data-sources/cbr_backup_ids_v3#parent_id DataOpentelekomcloudCbrBackupIdsV3#parent_id}.
	ParentId *string `field:"optional" json:"parentId" yaml:"parentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/data-sources/cbr_backup_ids_v3#resource_az DataOpentelekomcloudCbrBackupIdsV3#resource_az}.
	ResourceAz *string `field:"optional" json:"resourceAz" yaml:"resourceAz"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/data-sources/cbr_backup_ids_v3#resource_id DataOpentelekomcloudCbrBackupIdsV3#resource_id}.
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/data-sources/cbr_backup_ids_v3#resource_name DataOpentelekomcloudCbrBackupIdsV3#resource_name}.
	ResourceName *string `field:"optional" json:"resourceName" yaml:"resourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/data-sources/cbr_backup_ids_v3#resource_type DataOpentelekomcloudCbrBackupIdsV3#resource_type}.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/data-sources/cbr_backup_ids_v3#status DataOpentelekomcloudCbrBackupIdsV3#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/data-sources/cbr_backup_ids_v3#vault_id DataOpentelekomcloudCbrBackupIdsV3#vault_id}.
	VaultId *string `field:"optional" json:"vaultId" yaml:"vaultId"`
}

