// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dwsclusterv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DwsClusterV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#name DwsClusterV1#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#network_id DwsClusterV1#network_id}.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#node_type DwsClusterV1#node_type}.
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#number_of_node DwsClusterV1#number_of_node}.
	NumberOfNode *float64 `field:"required" json:"numberOfNode" yaml:"numberOfNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#security_group_id DwsClusterV1#security_group_id}.
	SecurityGroupId *string `field:"required" json:"securityGroupId" yaml:"securityGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#user_name DwsClusterV1#user_name}.
	UserName *string `field:"required" json:"userName" yaml:"userName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#user_pwd DwsClusterV1#user_pwd}.
	UserPwd *string `field:"required" json:"userPwd" yaml:"userPwd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#vpc_id DwsClusterV1#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#availability_zone DwsClusterV1#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// endpoints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#endpoints DwsClusterV1#endpoints}
	Endpoints interface{} `field:"optional" json:"endpoints" yaml:"endpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#id DwsClusterV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#keep_last_manual_snapshot DwsClusterV1#keep_last_manual_snapshot}.
	KeepLastManualSnapshot *float64 `field:"optional" json:"keepLastManualSnapshot" yaml:"keepLastManualSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#number_of_cn DwsClusterV1#number_of_cn}.
	NumberOfCn *float64 `field:"optional" json:"numberOfCn" yaml:"numberOfCn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#port DwsClusterV1#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// public_endpoints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#public_endpoints DwsClusterV1#public_endpoints}
	PublicEndpoints interface{} `field:"optional" json:"publicEndpoints" yaml:"publicEndpoints"`
	// public_ip block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#public_ip DwsClusterV1#public_ip}
	PublicIp *DwsClusterV1PublicIp `field:"optional" json:"publicIp" yaml:"publicIp"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/resources/dws_cluster_v1#timeouts DwsClusterV1#timeouts}
	Timeouts *DwsClusterV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

