// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cbrvaultv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CbrVaultV3Config struct {
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
	// billing block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/cbr_vault_v3#billing CbrVaultV3#billing}
	Billing *CbrVaultV3Billing `field:"required" json:"billing" yaml:"billing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/cbr_vault_v3#name CbrVaultV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/cbr_vault_v3#auto_bind CbrVaultV3#auto_bind}.
	AutoBind interface{} `field:"optional" json:"autoBind" yaml:"autoBind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/cbr_vault_v3#auto_expand CbrVaultV3#auto_expand}.
	AutoExpand interface{} `field:"optional" json:"autoExpand" yaml:"autoExpand"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/cbr_vault_v3#backup_policy_id CbrVaultV3#backup_policy_id}.
	BackupPolicyId *string `field:"optional" json:"backupPolicyId" yaml:"backupPolicyId"`
	// bind_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/cbr_vault_v3#bind_rules CbrVaultV3#bind_rules}
	BindRules interface{} `field:"optional" json:"bindRules" yaml:"bindRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/cbr_vault_v3#description CbrVaultV3#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/cbr_vault_v3#enterprise_project_id CbrVaultV3#enterprise_project_id}.
	EnterpriseProjectId *string `field:"optional" json:"enterpriseProjectId" yaml:"enterpriseProjectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/cbr_vault_v3#id CbrVaultV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/cbr_vault_v3#resource CbrVaultV3#resource}.
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.14/docs/resources/cbr_vault_v3#tags CbrVaultV3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

