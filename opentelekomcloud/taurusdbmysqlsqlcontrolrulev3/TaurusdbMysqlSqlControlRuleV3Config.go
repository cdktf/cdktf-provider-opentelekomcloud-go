// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package taurusdbmysqlsqlcontrolrulev3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TaurusdbMysqlSqlControlRuleV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/taurusdb_mysql_sql_control_rule_v3#instance_id TaurusdbMysqlSqlControlRuleV3#instance_id}.
	InstanceId *string `field:"required" json:"instanceId" yaml:"instanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/taurusdb_mysql_sql_control_rule_v3#max_concurrency TaurusdbMysqlSqlControlRuleV3#max_concurrency}.
	MaxConcurrency *float64 `field:"required" json:"maxConcurrency" yaml:"maxConcurrency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/taurusdb_mysql_sql_control_rule_v3#node_id TaurusdbMysqlSqlControlRuleV3#node_id}.
	NodeId *string `field:"required" json:"nodeId" yaml:"nodeId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/taurusdb_mysql_sql_control_rule_v3#pattern TaurusdbMysqlSqlControlRuleV3#pattern}.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/taurusdb_mysql_sql_control_rule_v3#sql_type TaurusdbMysqlSqlControlRuleV3#sql_type}.
	SqlType *string `field:"required" json:"sqlType" yaml:"sqlType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/taurusdb_mysql_sql_control_rule_v3#id TaurusdbMysqlSqlControlRuleV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.50/docs/resources/taurusdb_mysql_sql_control_rule_v3#timeouts TaurusdbMysqlSqlControlRuleV3#timeouts}
	Timeouts *TaurusdbMysqlSqlControlRuleV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

