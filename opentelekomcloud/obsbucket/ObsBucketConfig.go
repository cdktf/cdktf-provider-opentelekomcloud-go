// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucket

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObsBucketConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#bucket ObsBucket#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#acl ObsBucket#acl}.
	Acl *string `field:"optional" json:"acl" yaml:"acl"`
	// cors_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#cors_rule ObsBucket#cors_rule}
	CorsRule interface{} `field:"optional" json:"corsRule" yaml:"corsRule"`
	// event_notifications block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#event_notifications ObsBucket#event_notifications}
	EventNotifications interface{} `field:"optional" json:"eventNotifications" yaml:"eventNotifications"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#force_destroy ObsBucket#force_destroy}.
	ForceDestroy interface{} `field:"optional" json:"forceDestroy" yaml:"forceDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#id ObsBucket#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lifecycle_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#lifecycle_rule ObsBucket#lifecycle_rule}
	LifecycleRule interface{} `field:"optional" json:"lifecycleRule" yaml:"lifecycleRule"`
	// logging block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#logging ObsBucket#logging}
	Logging interface{} `field:"optional" json:"logging" yaml:"logging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#parallel_fs ObsBucket#parallel_fs}.
	ParallelFs interface{} `field:"optional" json:"parallelFs" yaml:"parallelFs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#region ObsBucket#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// server_side_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#server_side_encryption ObsBucket#server_side_encryption}
	ServerSideEncryption *ObsBucketServerSideEncryption `field:"optional" json:"serverSideEncryption" yaml:"serverSideEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#storage_class ObsBucket#storage_class}.
	StorageClass *string `field:"optional" json:"storageClass" yaml:"storageClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#tags ObsBucket#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#user_domain_names ObsBucket#user_domain_names}.
	UserDomainNames *[]*string `field:"optional" json:"userDomainNames" yaml:"userDomainNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#versioning ObsBucket#versioning}.
	Versioning interface{} `field:"optional" json:"versioning" yaml:"versioning"`
	// website block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#website ObsBucket#website}
	Website *ObsBucketWebsite `field:"optional" json:"website" yaml:"website"`
	// worm_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.47/docs/resources/obs_bucket#worm_policy ObsBucket#worm_policy}
	WormPolicy *ObsBucketWormPolicy `field:"optional" json:"wormPolicy" yaml:"wormPolicy"`
}

