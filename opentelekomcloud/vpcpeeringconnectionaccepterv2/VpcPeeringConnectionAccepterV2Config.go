package vpcpeeringconnectionaccepterv2

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpcPeeringConnectionAccepterV2Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.0/docs/resources/vpc_peering_connection_accepter_v2#vpc_peering_connection_id VpcPeeringConnectionAccepterV2#vpc_peering_connection_id}.
	VpcPeeringConnectionId *string `field:"required" json:"vpcPeeringConnectionId" yaml:"vpcPeeringConnectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.0/docs/resources/vpc_peering_connection_accepter_v2#accept VpcPeeringConnectionAccepterV2#accept}.
	Accept interface{} `field:"optional" json:"accept" yaml:"accept"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.0/docs/resources/vpc_peering_connection_accepter_v2#id VpcPeeringConnectionAccepterV2#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.0/docs/resources/vpc_peering_connection_accepter_v2#region VpcPeeringConnectionAccepterV2#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.0/docs/resources/vpc_peering_connection_accepter_v2#timeouts VpcPeeringConnectionAccepterV2#timeouts}
	Timeouts *VpcPeeringConnectionAccepterV2Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

