package cceaddonv3


type CceAddonV3Values struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/cce_addon_v3#basic CceAddonV3#basic}.
	Basic *map[string]*string `field:"required" json:"basic" yaml:"basic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/opentelekomcloud/opentelekomcloud/1.34.1/docs/resources/cce_addon_v3#custom CceAddonV3#custom}.
	Custom *map[string]*string `field:"required" json:"custom" yaml:"custom"`
}

