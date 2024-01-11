// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mrsclusterv1

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MrsClusterV1Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#available_zone_id MrsClusterV1#available_zone_id}.
	AvailableZoneId *string `field:"required" json:"availableZoneId" yaml:"availableZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#billing_type MrsClusterV1#billing_type}.
	BillingType *float64 `field:"required" json:"billingType" yaml:"billingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#cluster_name MrsClusterV1#cluster_name}.
	ClusterName *string `field:"required" json:"clusterName" yaml:"clusterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#cluster_version MrsClusterV1#cluster_version}.
	ClusterVersion *string `field:"required" json:"clusterVersion" yaml:"clusterVersion"`
	// component_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#component_list MrsClusterV1#component_list}
	ComponentList interface{} `field:"required" json:"componentList" yaml:"componentList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#core_node_num MrsClusterV1#core_node_num}.
	CoreNodeNum *float64 `field:"required" json:"coreNodeNum" yaml:"coreNodeNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#core_node_size MrsClusterV1#core_node_size}.
	CoreNodeSize *string `field:"required" json:"coreNodeSize" yaml:"coreNodeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#master_node_num MrsClusterV1#master_node_num}.
	MasterNodeNum *float64 `field:"required" json:"masterNodeNum" yaml:"masterNodeNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#master_node_size MrsClusterV1#master_node_size}.
	MasterNodeSize *string `field:"required" json:"masterNodeSize" yaml:"masterNodeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#node_public_cert_name MrsClusterV1#node_public_cert_name}.
	NodePublicCertName *string `field:"required" json:"nodePublicCertName" yaml:"nodePublicCertName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#safe_mode MrsClusterV1#safe_mode}.
	SafeMode *float64 `field:"required" json:"safeMode" yaml:"safeMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#subnet_id MrsClusterV1#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#vpc_id MrsClusterV1#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// add_jobs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#add_jobs MrsClusterV1#add_jobs}
	AddJobs interface{} `field:"optional" json:"addJobs" yaml:"addJobs"`
	// bootstrap_scripts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#bootstrap_scripts MrsClusterV1#bootstrap_scripts}
	BootstrapScripts interface{} `field:"optional" json:"bootstrapScripts" yaml:"bootstrapScripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#cluster_admin_secret MrsClusterV1#cluster_admin_secret}.
	ClusterAdminSecret *string `field:"optional" json:"clusterAdminSecret" yaml:"clusterAdminSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#cluster_type MrsClusterV1#cluster_type}.
	ClusterType *float64 `field:"optional" json:"clusterType" yaml:"clusterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#core_data_volume_count MrsClusterV1#core_data_volume_count}.
	CoreDataVolumeCount *float64 `field:"optional" json:"coreDataVolumeCount" yaml:"coreDataVolumeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#core_data_volume_size MrsClusterV1#core_data_volume_size}.
	CoreDataVolumeSize *float64 `field:"optional" json:"coreDataVolumeSize" yaml:"coreDataVolumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#core_data_volume_type MrsClusterV1#core_data_volume_type}.
	CoreDataVolumeType *string `field:"optional" json:"coreDataVolumeType" yaml:"coreDataVolumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#id MrsClusterV1#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#log_collection MrsClusterV1#log_collection}.
	LogCollection *float64 `field:"optional" json:"logCollection" yaml:"logCollection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#master_data_volume_count MrsClusterV1#master_data_volume_count}.
	MasterDataVolumeCount *float64 `field:"optional" json:"masterDataVolumeCount" yaml:"masterDataVolumeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#master_data_volume_size MrsClusterV1#master_data_volume_size}.
	MasterDataVolumeSize *float64 `field:"optional" json:"masterDataVolumeSize" yaml:"masterDataVolumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#master_data_volume_type MrsClusterV1#master_data_volume_type}.
	MasterDataVolumeType *string `field:"optional" json:"masterDataVolumeType" yaml:"masterDataVolumeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#region MrsClusterV1#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#tags MrsClusterV1#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#timeouts MrsClusterV1#timeouts}
	Timeouts *MrsClusterV1Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#volume_size MrsClusterV1#volume_size}.
	VolumeSize *float64 `field:"optional" json:"volumeSize" yaml:"volumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.16/docs/resources/mrs_cluster_v1#volume_type MrsClusterV1#volume_type}.
	VolumeType *string `field:"optional" json:"volumeType" yaml:"volumeType"`
}

