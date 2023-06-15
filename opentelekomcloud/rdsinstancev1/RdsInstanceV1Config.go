package rdsinstancev1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RdsInstanceV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#availabilityzone RdsInstanceV1#availabilityzone}.
	Availabilityzone *string `field:"required" json:"availabilityzone" yaml:"availabilityzone"`
	// datastore block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#datastore RdsInstanceV1#datastore}
	Datastore *RdsInstanceV1Datastore `field:"required" json:"datastore" yaml:"datastore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#dbrtpd RdsInstanceV1#dbrtpd}.
	Dbrtpd *string `field:"required" json:"dbrtpd" yaml:"dbrtpd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#flavorref RdsInstanceV1#flavorref}.
	Flavorref *string `field:"required" json:"flavorref" yaml:"flavorref"`
	// nics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#nics RdsInstanceV1#nics}
	Nics *RdsInstanceV1Nics `field:"required" json:"nics" yaml:"nics"`
	// securitygroup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#securitygroup RdsInstanceV1#securitygroup}
	Securitygroup *RdsInstanceV1Securitygroup `field:"required" json:"securitygroup" yaml:"securitygroup"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#volume RdsInstanceV1#volume}
	Volume *RdsInstanceV1Volume `field:"required" json:"volume" yaml:"volume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#vpc RdsInstanceV1#vpc}.
	Vpc *string `field:"required" json:"vpc" yaml:"vpc"`
	// backupstrategy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#backupstrategy RdsInstanceV1#backupstrategy}
	Backupstrategy *RdsInstanceV1Backupstrategy `field:"optional" json:"backupstrategy" yaml:"backupstrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#dbport RdsInstanceV1#dbport}.
	Dbport *string `field:"optional" json:"dbport" yaml:"dbport"`
	// ha block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#ha RdsInstanceV1#ha}
	Ha *RdsInstanceV1Ha `field:"optional" json:"ha" yaml:"ha"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#id RdsInstanceV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#name RdsInstanceV1#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#region RdsInstanceV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#tag RdsInstanceV1#tag}.
	Tag *map[string]*string `field:"optional" json:"tag" yaml:"tag"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.1/docs/resources/rds_instance_v1#timeouts RdsInstanceV1#timeouts}
	Timeouts *RdsInstanceV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

