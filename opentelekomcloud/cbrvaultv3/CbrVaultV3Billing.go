// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cbrvaultv3


type CbrVaultV3Billing struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cbr_vault_v3#object_type CbrVaultV3#object_type}.
	ObjectType *string `field:"required" json:"objectType" yaml:"objectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cbr_vault_v3#protect_type CbrVaultV3#protect_type}.
	ProtectType *string `field:"required" json:"protectType" yaml:"protectType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cbr_vault_v3#size CbrVaultV3#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cbr_vault_v3#charging_mode CbrVaultV3#charging_mode}.
	ChargingMode *string `field:"optional" json:"chargingMode" yaml:"chargingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cbr_vault_v3#cloud_type CbrVaultV3#cloud_type}.
	CloudType *string `field:"optional" json:"cloudType" yaml:"cloudType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cbr_vault_v3#consistent_level CbrVaultV3#consistent_level}.
	ConsistentLevel *string `field:"optional" json:"consistentLevel" yaml:"consistentLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cbr_vault_v3#console_url CbrVaultV3#console_url}.
	ConsoleUrl *string `field:"optional" json:"consoleUrl" yaml:"consoleUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cbr_vault_v3#extra_info CbrVaultV3#extra_info}.
	ExtraInfo *map[string]*string `field:"optional" json:"extraInfo" yaml:"extraInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cbr_vault_v3#is_auto_pay CbrVaultV3#is_auto_pay}.
	IsAutoPay interface{} `field:"optional" json:"isAutoPay" yaml:"isAutoPay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cbr_vault_v3#is_auto_renew CbrVaultV3#is_auto_renew}.
	IsAutoRenew interface{} `field:"optional" json:"isAutoRenew" yaml:"isAutoRenew"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cbr_vault_v3#period_num CbrVaultV3#period_num}.
	PeriodNum *float64 `field:"optional" json:"periodNum" yaml:"periodNum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.53/docs/resources/cbr_vault_v3#period_type CbrVaultV3#period_type}.
	PeriodType *string `field:"optional" json:"periodType" yaml:"periodType"`
}

