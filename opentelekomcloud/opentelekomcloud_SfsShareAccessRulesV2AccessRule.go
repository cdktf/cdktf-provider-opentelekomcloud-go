// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type SfsShareAccessRulesV2AccessRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/sfs_share_access_rules_v2#access_level SfsShareAccessRulesV2#access_level}.
	AccessLevel *string `field:"required" json:"accessLevel" yaml:"accessLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/sfs_share_access_rules_v2#access_to SfsShareAccessRulesV2#access_to}.
	AccessTo *string `field:"required" json:"accessTo" yaml:"accessTo"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/sfs_share_access_rules_v2#access_type SfsShareAccessRulesV2#access_type}.
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
}

