// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package taurusdbmysqlproxyv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TaurusdbMysqlProxyV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/taurusdb_mysql_proxy_v3#flavor TaurusdbMysqlProxyV3#flavor}.
	Flavor *string `field:"required" json:"flavor" yaml:"flavor"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/taurusdb_mysql_proxy_v3#instance_id TaurusdbMysqlProxyV3#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/taurusdb_mysql_proxy_v3#node_num TaurusdbMysqlProxyV3#node_num}.
	NodeNum *float64 `field:"required" json:"nodeNum" yaml:"nodeNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/taurusdb_mysql_proxy_v3#id TaurusdbMysqlProxyV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// master_node_weight block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/taurusdb_mysql_proxy_v3#master_node_weight TaurusdbMysqlProxyV3#master_node_weight}
	MasterNodeWeight *TaurusdbMysqlProxyV3MasterNodeWeight `field:"optional" json:"masterNodeWeight" yaml:"masterNodeWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/taurusdb_mysql_proxy_v3#proxy_mode TaurusdbMysqlProxyV3#proxy_mode}.
	ProxyMode *string `field:"optional" json:"proxyMode" yaml:"proxyMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/taurusdb_mysql_proxy_v3#proxy_name TaurusdbMysqlProxyV3#proxy_name}.
	ProxyName *string `field:"optional" json:"proxyName" yaml:"proxyName"`
	// readonly_nodes_weight block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/taurusdb_mysql_proxy_v3#readonly_nodes_weight TaurusdbMysqlProxyV3#readonly_nodes_weight}
	ReadonlyNodesWeight interface{} `field:"optional" json:"readonlyNodesWeight" yaml:"readonlyNodesWeight"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/taurusdb_mysql_proxy_v3#timeouts TaurusdbMysqlProxyV3#timeouts}
	Timeouts *TaurusdbMysqlProxyV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

