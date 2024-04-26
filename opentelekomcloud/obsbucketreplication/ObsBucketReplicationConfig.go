// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package obsbucketreplication

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ObsBucketReplicationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/obs_bucket_replication#agency ObsBucketReplication#agency}.
	Agency *string `field:"required" json:"agency" yaml:"agency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/obs_bucket_replication#bucket ObsBucketReplication#bucket}.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/obs_bucket_replication#destination_bucket ObsBucketReplication#destination_bucket}.
	DestinationBucket *string `field:"required" json:"destinationBucket" yaml:"destinationBucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/obs_bucket_replication#id ObsBucketReplication#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/obs_bucket_replication#region ObsBucketReplication#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.7/docs/resources/obs_bucket_replication#rule ObsBucketReplication#rule}
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
}

