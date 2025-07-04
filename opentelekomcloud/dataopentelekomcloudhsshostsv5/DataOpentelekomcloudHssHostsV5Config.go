// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataopentelekomcloudhsshostsv5

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataOpentelekomcloudHssHostsV5Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/data-sources/hss_hosts_v5#agent_status DataOpentelekomcloudHssHostsV5#agent_status}.
	AgentStatus *string `field:"optional" json:"agentStatus" yaml:"agentStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/data-sources/hss_hosts_v5#asset_value DataOpentelekomcloudHssHostsV5#asset_value}.
	AssetValue *string `field:"optional" json:"assetValue" yaml:"assetValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/data-sources/hss_hosts_v5#detect_result DataOpentelekomcloudHssHostsV5#detect_result}.
	DetectResult *string `field:"optional" json:"detectResult" yaml:"detectResult"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/data-sources/hss_hosts_v5#group_id DataOpentelekomcloudHssHostsV5#group_id}.
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/data-sources/hss_hosts_v5#host_id DataOpentelekomcloudHssHostsV5#host_id}.
	HostId *string `field:"optional" json:"hostId" yaml:"hostId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/data-sources/hss_hosts_v5#id DataOpentelekomcloudHssHostsV5#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/data-sources/hss_hosts_v5#name DataOpentelekomcloudHssHostsV5#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/data-sources/hss_hosts_v5#os_type DataOpentelekomcloudHssHostsV5#os_type}.
	OsType *string `field:"optional" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/data-sources/hss_hosts_v5#policy_group_id DataOpentelekomcloudHssHostsV5#policy_group_id}.
	PolicyGroupId *string `field:"optional" json:"policyGroupId" yaml:"policyGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/data-sources/hss_hosts_v5#protect_charging_mode DataOpentelekomcloudHssHostsV5#protect_charging_mode}.
	ProtectChargingMode *string `field:"optional" json:"protectChargingMode" yaml:"protectChargingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/data-sources/hss_hosts_v5#protect_status DataOpentelekomcloudHssHostsV5#protect_status}.
	ProtectStatus *string `field:"optional" json:"protectStatus" yaml:"protectStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/data-sources/hss_hosts_v5#protect_version DataOpentelekomcloudHssHostsV5#protect_version}.
	ProtectVersion *string `field:"optional" json:"protectVersion" yaml:"protectVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.42/docs/data-sources/hss_hosts_v5#status DataOpentelekomcloudHssHostsV5#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

