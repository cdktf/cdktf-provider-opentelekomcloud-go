// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package taurusdbmysqlproxyv3


type TaurusdbMysqlProxyV3ReadonlyNodesWeight struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_proxy_v3#id TaurusdbMysqlProxyV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.52/docs/resources/taurusdb_mysql_proxy_v3#weight TaurusdbMysqlProxyV3#weight}.
	Weight *float64 `field:"required" json:"weight" yaml:"weight"`
}

