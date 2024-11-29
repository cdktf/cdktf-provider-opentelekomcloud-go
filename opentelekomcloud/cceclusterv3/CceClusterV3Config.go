// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cceclusterv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CceClusterV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#cluster_type CceClusterV3#cluster_type}.
	ClusterType *string `field:"required" json:"clusterType" yaml:"clusterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#container_network_type CceClusterV3#container_network_type}.
	ContainerNetworkType *string `field:"required" json:"containerNetworkType" yaml:"containerNetworkType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#flavor_id CceClusterV3#flavor_id}.
	FlavorId *string `field:"required" json:"flavorId" yaml:"flavorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#name CceClusterV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#subnet_id CceClusterV3#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#vpc_id CceClusterV3#vpc_id}.
	VpcId *string `field:"required" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#annotations CceClusterV3#annotations}.
	Annotations *map[string]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// authenticating_proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#authenticating_proxy CceClusterV3#authenticating_proxy}
	AuthenticatingProxy *CceClusterV3AuthenticatingProxy `field:"optional" json:"authenticatingProxy" yaml:"authenticatingProxy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#authenticating_proxy_ca CceClusterV3#authenticating_proxy_ca}.
	AuthenticatingProxyCa *string `field:"optional" json:"authenticatingProxyCa" yaml:"authenticatingProxyCa"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#authentication_mode CceClusterV3#authentication_mode}.
	AuthenticationMode *string `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#billing_mode CceClusterV3#billing_mode}.
	BillingMode *float64 `field:"optional" json:"billingMode" yaml:"billingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#cluster_version CceClusterV3#cluster_version}.
	ClusterVersion *string `field:"optional" json:"clusterVersion" yaml:"clusterVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#container_network_cidr CceClusterV3#container_network_cidr}.
	ContainerNetworkCidr *string `field:"optional" json:"containerNetworkCidr" yaml:"containerNetworkCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#delete_all_network CceClusterV3#delete_all_network}.
	DeleteAllNetwork *string `field:"optional" json:"deleteAllNetwork" yaml:"deleteAllNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#delete_all_storage CceClusterV3#delete_all_storage}.
	DeleteAllStorage *string `field:"optional" json:"deleteAllStorage" yaml:"deleteAllStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#delete_efs CceClusterV3#delete_efs}.
	DeleteEfs *string `field:"optional" json:"deleteEfs" yaml:"deleteEfs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#delete_eni CceClusterV3#delete_eni}.
	DeleteEni *string `field:"optional" json:"deleteEni" yaml:"deleteEni"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#delete_evs CceClusterV3#delete_evs}.
	DeleteEvs *string `field:"optional" json:"deleteEvs" yaml:"deleteEvs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#delete_net CceClusterV3#delete_net}.
	DeleteNet *string `field:"optional" json:"deleteNet" yaml:"deleteNet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#delete_obs CceClusterV3#delete_obs}.
	DeleteObs *string `field:"optional" json:"deleteObs" yaml:"deleteObs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#delete_sfs CceClusterV3#delete_sfs}.
	DeleteSfs *string `field:"optional" json:"deleteSfs" yaml:"deleteSfs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#description CceClusterV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#eip CceClusterV3#eip}.
	Eip *string `field:"optional" json:"eip" yaml:"eip"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#enable_volume_encryption CceClusterV3#enable_volume_encryption}.
	EnableVolumeEncryption interface{} `field:"optional" json:"enableVolumeEncryption" yaml:"enableVolumeEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#eni_subnet_cidr CceClusterV3#eni_subnet_cidr}.
	EniSubnetCidr *string `field:"optional" json:"eniSubnetCidr" yaml:"eniSubnetCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#eni_subnet_id CceClusterV3#eni_subnet_id}.
	EniSubnetId *string `field:"optional" json:"eniSubnetId" yaml:"eniSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#extend_param CceClusterV3#extend_param}.
	ExtendParam *map[string]*string `field:"optional" json:"extendParam" yaml:"extendParam"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#highway_subnet_id CceClusterV3#highway_subnet_id}.
	HighwaySubnetId *string `field:"optional" json:"highwaySubnetId" yaml:"highwaySubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#id CceClusterV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#ignore_addons CceClusterV3#ignore_addons}.
	IgnoreAddons interface{} `field:"optional" json:"ignoreAddons" yaml:"ignoreAddons"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#ignore_certificate_clusters_data CceClusterV3#ignore_certificate_clusters_data}.
	IgnoreCertificateClustersData interface{} `field:"optional" json:"ignoreCertificateClustersData" yaml:"ignoreCertificateClustersData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#ignore_certificate_users_data CceClusterV3#ignore_certificate_users_data}.
	IgnoreCertificateUsersData interface{} `field:"optional" json:"ignoreCertificateUsersData" yaml:"ignoreCertificateUsersData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#kube_proxy_mode CceClusterV3#kube_proxy_mode}.
	KubeProxyMode *string `field:"optional" json:"kubeProxyMode" yaml:"kubeProxyMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#kubernetes_svc_ip_range CceClusterV3#kubernetes_svc_ip_range}.
	KubernetesSvcIpRange *string `field:"optional" json:"kubernetesSvcIpRange" yaml:"kubernetesSvcIpRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#labels CceClusterV3#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#multi_az CceClusterV3#multi_az}.
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#no_addons CceClusterV3#no_addons}.
	NoAddons interface{} `field:"optional" json:"noAddons" yaml:"noAddons"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#region CceClusterV3#region}.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#security_group_id CceClusterV3#security_group_id}.
	SecurityGroupId *string `field:"optional" json:"securityGroupId" yaml:"securityGroupId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.26/docs/resources/cce_cluster_v3#timeouts CceClusterV3#timeouts}
	Timeouts *CceClusterV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

