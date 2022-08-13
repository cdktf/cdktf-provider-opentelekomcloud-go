// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type CceAddonV3Values struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_addon_v3#basic CceAddonV3#basic}.
	Basic *map[string]*string `field:"required" json:"basic" yaml:"basic"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/cce_addon_v3#custom CceAddonV3#custom}.
	Custom *map[string]*string `field:"required" json:"custom" yaml:"custom"`
}

