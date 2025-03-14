// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodepoolv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CceNodePoolV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#cluster_id CceNodePoolV3#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// data_volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#data_volumes CceNodePoolV3#data_volumes}
	DataVolumes interface{} `field:"required" json:"dataVolumes" yaml:"dataVolumes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#flavor CceNodePoolV3#flavor}.
	Flavor *string `field:"required" json:"flavor" yaml:"flavor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#initial_node_count CceNodePoolV3#initial_node_count}.
	InitialNodeCount *float64 `field:"required" json:"initialNodeCount" yaml:"initialNodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#name CceNodePoolV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// root_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#root_volume CceNodePoolV3#root_volume}
	RootVolume *CceNodePoolV3RootVolume `field:"required" json:"rootVolume" yaml:"rootVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#agency_name CceNodePoolV3#agency_name}.
	AgencyName *string `field:"optional" json:"agencyName" yaml:"agencyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#availability_zone CceNodePoolV3#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#docker_base_size CceNodePoolV3#docker_base_size}.
	DockerBaseSize *float64 `field:"optional" json:"dockerBaseSize" yaml:"dockerBaseSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#docker_lvm_config_override CceNodePoolV3#docker_lvm_config_override}.
	DockerLvmConfigOverride *string `field:"optional" json:"dockerLvmConfigOverride" yaml:"dockerLvmConfigOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#id CceNodePoolV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#k8s_tags CceNodePoolV3#k8s_tags}.
	K8STags *map[string]*string `field:"optional" json:"k8STags" yaml:"k8STags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#key_pair CceNodePoolV3#key_pair}.
	KeyPair *string `field:"optional" json:"keyPair" yaml:"keyPair"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#max_node_count CceNodePoolV3#max_node_count}.
	MaxNodeCount *float64 `field:"optional" json:"maxNodeCount" yaml:"maxNodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#max_pods CceNodePoolV3#max_pods}.
	MaxPods *float64 `field:"optional" json:"maxPods" yaml:"maxPods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#min_node_count CceNodePoolV3#min_node_count}.
	MinNodeCount *float64 `field:"optional" json:"minNodeCount" yaml:"minNodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#os CceNodePoolV3#os}.
	Os *string `field:"optional" json:"os" yaml:"os"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#password CceNodePoolV3#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#postinstall CceNodePoolV3#postinstall}.
	Postinstall *string `field:"optional" json:"postinstall" yaml:"postinstall"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#preinstall CceNodePoolV3#preinstall}.
	Preinstall *string `field:"optional" json:"preinstall" yaml:"preinstall"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#priority CceNodePoolV3#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#runtime CceNodePoolV3#runtime}.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#scale_down_cooldown_time CceNodePoolV3#scale_down_cooldown_time}.
	ScaleDownCooldownTime *float64 `field:"optional" json:"scaleDownCooldownTime" yaml:"scaleDownCooldownTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#scale_enable CceNodePoolV3#scale_enable}.
	ScaleEnable interface{} `field:"optional" json:"scaleEnable" yaml:"scaleEnable"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#security_group_ids CceNodePoolV3#security_group_ids}.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#server_group_reference CceNodePoolV3#server_group_reference}.
	ServerGroupReference *string `field:"optional" json:"serverGroupReference" yaml:"serverGroupReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#storage CceNodePoolV3#storage}.
	Storage *string `field:"optional" json:"storage" yaml:"storage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#subnet_id CceNodePoolV3#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// taints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#taints CceNodePoolV3#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#timeouts CceNodePoolV3#timeouts}
	Timeouts *CceNodePoolV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.33/docs/resources/cce_node_pool_v3#user_tags CceNodePoolV3#user_tags}.
	UserTags *map[string]*string `field:"optional" json:"userTags" yaml:"userTags"`
}

