// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ccenodev3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CceNodeV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#availability_zone CceNodeV3#availability_zone}.
	AvailabilityZone *string `field:"required" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#cluster_id CceNodeV3#cluster_id}.
	ClusterId *string `field:"required" json:"clusterId" yaml:"clusterId"`
	// data_volumes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#data_volumes CceNodeV3#data_volumes}
	DataVolumes interface{} `field:"required" json:"dataVolumes" yaml:"dataVolumes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#flavor_id CceNodeV3#flavor_id}.
	FlavorId *string `field:"required" json:"flavorId" yaml:"flavorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#key_pair CceNodeV3#key_pair}.
	KeyPair *string `field:"required" json:"keyPair" yaml:"keyPair"`
	// root_volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#root_volume CceNodeV3#root_volume}
	RootVolume *CceNodeV3RootVolume `field:"required" json:"rootVolume" yaml:"rootVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#agency_name CceNodeV3#agency_name}.
	AgencyName *string `field:"optional" json:"agencyName" yaml:"agencyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#annotations CceNodeV3#annotations}.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#bandwidth_charge_mode CceNodeV3#bandwidth_charge_mode}.
	BandwidthChargeMode *string `field:"optional" json:"bandwidthChargeMode" yaml:"bandwidthChargeMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#bandwidth_size CceNodeV3#bandwidth_size}.
	BandwidthSize *float64 `field:"optional" json:"bandwidthSize" yaml:"bandwidthSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#billing_mode CceNodeV3#billing_mode}.
	BillingMode *float64 `field:"optional" json:"billingMode" yaml:"billingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#dedicated_host_id CceNodeV3#dedicated_host_id}.
	DedicatedHostId *string `field:"optional" json:"dedicatedHostId" yaml:"dedicatedHostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#docker_base_size CceNodeV3#docker_base_size}.
	DockerBaseSize *float64 `field:"optional" json:"dockerBaseSize" yaml:"dockerBaseSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#docker_lvm_config_override CceNodeV3#docker_lvm_config_override}.
	DockerLvmConfigOverride *string `field:"optional" json:"dockerLvmConfigOverride" yaml:"dockerLvmConfigOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#ecs_performance_type CceNodeV3#ecs_performance_type}.
	EcsPerformanceType *string `field:"optional" json:"ecsPerformanceType" yaml:"ecsPerformanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#eip_count CceNodeV3#eip_count}.
	EipCount *float64 `field:"optional" json:"eipCount" yaml:"eipCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#eip_ids CceNodeV3#eip_ids}.
	EipIds *[]*string `field:"optional" json:"eipIds" yaml:"eipIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#extend_param_charging_mode CceNodeV3#extend_param_charging_mode}.
	ExtendParamChargingMode *float64 `field:"optional" json:"extendParamChargingMode" yaml:"extendParamChargingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#id CceNodeV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#iptype CceNodeV3#iptype}.
	Iptype *string `field:"optional" json:"iptype" yaml:"iptype"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#k8s_tags CceNodeV3#k8s_tags}.
	K8STags *map[string]*string `field:"optional" json:"k8STags" yaml:"k8STags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#labels CceNodeV3#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#max_pods CceNodeV3#max_pods}.
	MaxPods *float64 `field:"optional" json:"maxPods" yaml:"maxPods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#name CceNodeV3#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#order_id CceNodeV3#order_id}.
	OrderId *string `field:"optional" json:"orderId" yaml:"orderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#os CceNodeV3#os}.
	Os *string `field:"optional" json:"os" yaml:"os"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#postinstall CceNodeV3#postinstall}.
	Postinstall *string `field:"optional" json:"postinstall" yaml:"postinstall"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#preinstall CceNodeV3#preinstall}.
	Preinstall *string `field:"optional" json:"preinstall" yaml:"preinstall"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#private_ip CceNodeV3#private_ip}.
	PrivateIp *string `field:"optional" json:"privateIp" yaml:"privateIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#product_id CceNodeV3#product_id}.
	ProductId *string `field:"optional" json:"productId" yaml:"productId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#public_key CceNodeV3#public_key}.
	PublicKey *string `field:"optional" json:"publicKey" yaml:"publicKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#region CceNodeV3#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#runtime CceNodeV3#runtime}.
	Runtime *string `field:"optional" json:"runtime" yaml:"runtime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#sharetype CceNodeV3#sharetype}.
	Sharetype *string `field:"optional" json:"sharetype" yaml:"sharetype"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#subnet_id CceNodeV3#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#tags CceNodeV3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// taints block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#taints CceNodeV3#taints}
	Taints interface{} `field:"optional" json:"taints" yaml:"taints"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/cce_node_v3#timeouts CceNodeV3#timeouts}
	Timeouts *CceNodeV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

