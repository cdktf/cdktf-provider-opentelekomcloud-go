// Prebuilt opentelekomcloud Provider for Terraform CDK (cdktf)
package opentelekomcloud


type IdentityRoleV3Statement struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/identity_role_v3#action IdentityRoleV3#action}.
	Action *[]*string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/identity_role_v3#effect IdentityRoleV3#effect}.
	Effect *string `field:"required" json:"effect" yaml:"effect"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/opentelekomcloud/r/identity_role_v3#resource IdentityRoleV3#resource}.
	Resource *[]*string `field:"optional" json:"resource" yaml:"resource"`
}

