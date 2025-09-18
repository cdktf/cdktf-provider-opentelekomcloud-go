// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cbrvaultv3


type CbrVaultV3Resource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cbr_vault_v3#backup_count CbrVaultV3#backup_count}.
	BackupCount *float64 `field:"optional" json:"backupCount" yaml:"backupCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cbr_vault_v3#backup_size CbrVaultV3#backup_size}.
	BackupSize *float64 `field:"optional" json:"backupSize" yaml:"backupSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cbr_vault_v3#exclude_volumes CbrVaultV3#exclude_volumes}.
	ExcludeVolumes *[]*string `field:"optional" json:"excludeVolumes" yaml:"excludeVolumes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cbr_vault_v3#id CbrVaultV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cbr_vault_v3#include_volumes CbrVaultV3#include_volumes}.
	IncludeVolumes *[]*string `field:"optional" json:"includeVolumes" yaml:"includeVolumes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cbr_vault_v3#name CbrVaultV3#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cbr_vault_v3#protect_status CbrVaultV3#protect_status}.
	ProtectStatus *string `field:"optional" json:"protectStatus" yaml:"protectStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cbr_vault_v3#size CbrVaultV3#size}.
	Size *float64 `field:"optional" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.36.48/docs/resources/cbr_vault_v3#type CbrVaultV3#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

