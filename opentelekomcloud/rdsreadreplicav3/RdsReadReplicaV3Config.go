// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rdsreadreplicav3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsReadReplicaV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/rds_read_replica_v3#flavor_ref RdsReadReplicaV3#flavor_ref}.
	FlavorRef *string `field:"required" json:"flavorRef" yaml:"flavorRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/rds_read_replica_v3#name RdsReadReplicaV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/rds_read_replica_v3#replica_of_id RdsReadReplicaV3#replica_of_id}.
	ReplicaOfId *string `field:"required" json:"replicaOfId" yaml:"replicaOfId"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/rds_read_replica_v3#volume RdsReadReplicaV3#volume}
	Volume *RdsReadReplicaV3Volume `field:"required" json:"volume" yaml:"volume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/rds_read_replica_v3#availability_zone RdsReadReplicaV3#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/rds_read_replica_v3#id RdsReadReplicaV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/rds_read_replica_v3#public_ips RdsReadReplicaV3#public_ips}.
	PublicIps *[]*string `field:"optional" json:"publicIps" yaml:"publicIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/rds_read_replica_v3#region RdsReadReplicaV3#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.15/docs/resources/rds_read_replica_v3#timeouts RdsReadReplicaV3#timeouts}
	Timeouts *RdsReadReplicaV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

