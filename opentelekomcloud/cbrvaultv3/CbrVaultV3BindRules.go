package cbrvaultv3


type CbrVaultV3BindRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.2/docs/resources/cbr_vault_v3#key CbrVaultV3#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.35.2/docs/resources/cbr_vault_v3#value CbrVaultV3#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

